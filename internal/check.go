package internal

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"h12.io/socks"
)

const checkTimeout = 2 * time.Second

// CheckURLs 用于验证代理的 GET 访问地址，两个都需在 2 秒内成功
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

// CheckProxy 使用代理 GET 访问 CheckURLs 中的每个 URL，全部在 2 秒内成功则通过
func CheckProxy(p *Proxy) bool {
	if p == nil || !IsAllowedProtocol(p.Protocol) {
		return false
	}
	proxyURL := p.String()
	for _, targetURL := range CheckURLs {
		ok := checkOne(p, proxyURL, targetURL)
		if !ok {
			return false
		}
	}
	return true
}

func checkOne(p *Proxy, proxyURL, targetURL string) bool {
	client, err := httpClientViaProxy(p, proxyURL)
	if err != nil {
		return false
	}
	client.Timeout = checkTimeout
	ctx, cancel := context.WithTimeout(context.Background(), checkTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		return false
	}
	start := time.Now()
	resp, err := client.Do(req)
	elapsed := time.Since(start)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
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
