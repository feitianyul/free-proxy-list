package internal

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"h12.io/socks"
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
	checkTimeout      = 2 * time.Second
	defaultCheckWorkers = 500
	maxCheckWorkers   = 1000
)

// CheckWorkers 并发校验的 worker 数量，由 main 通过 flag/env 设置，未设置时用 GetCheckWorkers() 的默认值
var CheckWorkers int

// CheckURLs 用于验证代理的访问地址，两个都需在 2 秒内成功（使用 HEAD 减少传输）
var CheckURLs = []string{
	"https://www.eastmoney.com/",
	"https://sinajs.cn/",
}

// IsAllowedProtocol 是否为允许保留的代理类型：http、https、socks4、socks5
func IsAllowedProtocol(protocol string) bool {
	switch strings.ToLower(protocol) {
	case "http", "https", "socks4", "socks4a", "socks5", "socks5h":
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
	client, err := httpClientViaProxy(p, proxyURL)
	if err != nil {
		return false
	}
	client.Timeout = checkTimeout
	ctx, cancel := context.WithTimeout(context.Background(), checkTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, targetURL, nil)
	if err != nil {
		return false
	}
	start := time.Now()
	resp, err := client.Do(req)
	elapsed := time.Since(start)
	if err != nil {
		return false
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
			return false
		}
	}
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil || resp == nil || resp.StatusCode != http.StatusOK {
		return false
	}
	return elapsed <= checkTimeout
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
	case "socks4", "socks4a", "socks5", "socks5h":
		// h12.io/socks 支持 socks5、socks4、socks4a
		socksScheme := "socks5"
		if scheme == "socks5h" {
			socksScheme = "socks5"
		} else if scheme == "socks4" {
			socksScheme = "socks4"
		} else if scheme == "socks4a" {
			socksScheme = "socks4a"
		}
		addr := net.JoinHostPort(p.IP, strconv.Itoa(p.Port))
		uri := socksScheme + "://" + addr
		if p.User != "" {
			if p.Passwd != "" {
				uri = socksScheme + "://" + url.UserPassword(p.User, p.Passwd).String() + "@" + addr
			} else {
				uri = socksScheme + "://" + p.User + "@" + addr
			}
		}
		dial := socks.Dial(uri + "?timeout=" + checkTimeout.String())
		tr := &http.Transport{
			Dial:                  dial,
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
