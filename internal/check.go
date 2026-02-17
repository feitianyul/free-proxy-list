package internal

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func init() {
	// 环境变量 GFP_CHECK_WORKERS 可覆盖默认并发数（main 未传 -check-workers 时生效）
	if s := strings.TrimSpace(os.Getenv("GFP_CHECK_WORKERS")); s != "" {
		if n, err := strconv.Atoi(s); err == nil && n > 0 {
			CheckWorkers = n
		}
	}
}

const (
	checkTimeout        = 2 * time.Second
	defaultCheckWorkers = 4000
	maxCheckWorkers     = 4000
)

// CheckWorkers 并发校验的 worker 数量，由 main 通过 flag/env 设置，未设置时用 GetCheckWorkers() 的默认值
var CheckWorkers int

// CheckURLs 三域名验证（eastmoney + sse + sina 财经），均需在 2 秒内成功（使用 HEAD 减少传输）
var CheckURLs = []string{
	"https://www.eastmoney.com/",
	"https://www.sse.com.cn/",
	"https://finance.sina.com.cn/",
}

// IsAllowedProtocol 是否为允许保留的代理类型：仅 http、https
func IsAllowedProtocol(protocol string) bool {
	switch strings.ToLower(protocol) {
	case "http", "https":
		return true
	default:
		return false
	}
}

// CheckProxy 使用代理访问 CheckURLs 中的每个 URL（HEAD），全部在 2 秒内成功则通过；双 URL 并行
func CheckProxy(p *Proxy) bool {
	if p == nil || !IsAllowedProtocol(p.Protocol) {
		return false
	}
	proxyURL := p.String()
	var wg sync.WaitGroup
	results := make([]bool, len(CheckURLs))
	for i, targetURL := range CheckURLs {
		wg.Add(1)
		go func(idx int, target string) {
			defer wg.Done()
			results[idx] = checkOne(p, proxyURL, target)
		}(i, targetURL)
	}
	wg.Wait()
	for _, ok := range results {
		if !ok {
			return false
		}
	}
	return true
}

// checkOne 使用 HEAD 请求验证代理可达性，2 秒内返回 200 则通过；若 HEAD 返回 405 则回退 GET
func checkOne(p *Proxy, proxyURL, targetURL string) bool {
	ok, _, _ := checkOneWithResult(p, proxyURL, targetURL)
	return ok
}

// checkOneWithResult 同 checkOne，并返回耗时与失败时的简短错误信息（供双协议表格使用）
func checkOneWithResult(p *Proxy, proxyURL, targetURL string) (ok bool, elapsed time.Duration, errMsg string) {
	client, err := httpClientViaProxy(p, proxyURL)
	if err != nil {
		return false, 0, shortErr(err)
	}
	client.Timeout = checkTimeout
	ctx, cancel := context.WithTimeout(context.Background(), checkTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, targetURL, nil)
	if err != nil {
		return false, 0, shortErr(err)
	}
	start := time.Now()
	resp, err := client.Do(req)
	elapsed = time.Since(start)
	if err != nil {
		return false, elapsed, shortErr(err)
	}
	if resp.StatusCode == http.StatusMethodNotAllowed {
		resp.Body.Close()
		resp = nil
		ctx2, cancel2 := context.WithTimeout(context.Background(), checkTimeout)
		defer cancel2()
		req2, _ := http.NewRequestWithContext(ctx2, http.MethodGet, targetURL, nil)
		start = time.Now()
		resp, err = client.Do(req2)
		elapsed = time.Since(start)
		if err != nil {
			return false, elapsed, shortErr(err)
		}
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil || resp == nil {
		return false, elapsed, shortErr(err)
	}
	if resp.StatusCode != http.StatusOK {
		return false, elapsed, truncateStr(resp.Status, 40)
	}
	return elapsed <= checkTimeout, elapsed, ""
}

func shortErr(e error) string {
	if e == nil {
		return ""
	}
	return truncateStr(e.Error(), 40)
}

func truncateStr(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max]
}

func httpClientViaProxy(p *Proxy, proxyURL string) (*http.Client, error) {
	scheme := strings.ToLower(p.Protocol)
	switch scheme {
	case "http", "https":
		u, err := url.Parse(proxyURL)
		if err != nil {
			return nil, err
		}
		tr := &http.Transport{
			Proxy:                 http.ProxyURL(u),
			TLSHandshakeTimeout:   checkTimeout,
			ResponseHeaderTimeout: checkTimeout,
		}
		return &http.Client{Transport: tr}, nil
	default:
		return nil, errUnsupportedProtocol
	}
}

var errUnsupportedProtocol = errors.New("unsupported protocol for check")

// GetCheckWorkers 返回并发校验 worker 数，默认 defaultCheckWorkers，上限 maxCheckWorkers
func GetCheckWorkers() int {
	n := CheckWorkers
	if n <= 0 {
		n = defaultCheckWorkers
	}
	if n > maxCheckWorkers {
		n = maxCheckWorkers
	}
	return n
}

// CheckProxyAsHTTP 以 HTTP 代理方式校验三域名（eastmoney + sse + sina），每域 2s 内成功，返回是否通过、每域延迟、失败时的简短错误
func CheckProxyAsHTTP(p *Proxy) (ok bool, elapsed []time.Duration, errMsg string) {
	if p == nil {
		return false, nil, "nil proxy"
	}
	orig := p.Protocol
	p.Protocol = "http"
	proxyURL := p.String()
	p.Protocol = orig
	elapsed = make([]time.Duration, len(CheckURLs))
	for i, targetURL := range CheckURLs {
		oneOk, oneElapsed, oneErr := checkOneWithResult(p, proxyURL, targetURL)
		elapsed[i] = oneElapsed
		if !oneOk {
			return false, elapsed, oneErr
		}
	}
	return true, elapsed, ""
}

// CheckProxyAsHTTPS 以 HTTPS 代理方式校验三域名（eastmoney + sse + sina），每域 2s 内成功，返回是否通过、每域延迟、失败时的简短错误
func CheckProxyAsHTTPS(p *Proxy) (ok bool, elapsed []time.Duration, errMsg string) {
	if p == nil {
		return false, nil, "nil proxy"
	}
	orig := p.Protocol
	p.Protocol = "https"
	proxyURL := p.String()
	p.Protocol = orig
	elapsed = make([]time.Duration, len(CheckURLs))
	for i, targetURL := range CheckURLs {
		oneOk, oneElapsed, oneErr := checkOneWithResult(p, proxyURL, targetURL)
		elapsed[i] = oneElapsed
		if !oneOk {
			return false, elapsed, oneErr
		}
	}
	return true, elapsed, ""
}

// ValidateProxiesConcurrent 并发校验 proxies，通过者由单 goroutine 调用 Save，返回通过数量
func ValidateProxiesConcurrent(proxies []*Proxy, workers int) int {
	if workers <= 0 {
		workers = GetCheckWorkers()
	}
	if workers > maxCheckWorkers {
		workers = maxCheckWorkers
	}
	if len(proxies) == 0 {
		return 0
	}
	taskCh := make(chan *Proxy, len(proxies))
	resultCh := make(chan *Proxy, workers*2)
	for _, p := range proxies {
		taskCh <- p
	}
	close(taskCh)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range taskCh {
				if CheckProxy(p) {
					resultCh <- p
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var total int
	for p := range resultCh {
		Save(p)
		total++
	}
	return total
}

// ValidateProxiesDual 对每个候选代理分别以 HTTP 与 HTTPS 各测一次，结果写 db 并追加到双协议结果列表；返回至少通过一种协议的代理数量
func ValidateProxiesDual(proxies []*Proxy, workers int) int {
	if workers <= 0 {
		workers = GetCheckWorkers()
	}
	if workers > maxCheckWorkers {
		workers = maxCheckWorkers
	}
	if len(proxies) == 0 {
		return 0
	}
	taskCh := make(chan *Proxy, len(proxies))
	resultCh := make(chan *ProxyResult, workers*2)
	for _, p := range proxies {
		taskCh <- p
	}
	close(taskCh)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range taskCh {
				httpOk, httpElapsed, httpErr := CheckProxyAsHTTP(p)
				httpsOk, httpsElapsed, httpsErr := CheckProxyAsHTTPS(p)
				protocol := ""
				if httpOk && httpsOk {
					protocol = "http/s"
				} else if httpOk {
					protocol = "http"
				} else if httpsOk {
					protocol = "https"
				}
				r := &ProxyResult{
					IP: p.IP, Port: p.Port, User: p.User, Passwd: p.Passwd,
					HTTPOk: httpOk, HTTPElapsed: httpElapsed, HTTPErr: httpErr,
					HTTPSOk: httpsOk, HTTPSElapsed: httpsElapsed, HTTPSErr: httpsErr,
					Protocol: protocol,
				}
				resultCh <- r
			}
		}()
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var passed int
	for r := range resultCh {
		if r.HTTPOk {
			Save(r.Proxy("http"))
		}
		if r.HTTPSOk {
			Save(r.Proxy("https"))
		}
		if r.HTTPOk || r.HTTPSOk {
			passed++
		}
		AppendDualResult(r)
	}
	return passed
}
