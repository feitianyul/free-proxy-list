package internal

import (
	"strconv"
	"strings"
	"sync"
	"time"
)

// ProxyResult 保存单个代理的双协议（HTTP / HTTPS）校验结果，用于写 http+s.txt 与 README 表格
type ProxyResult struct {
	IP          string
	Port        int
	User        string
	Passwd      string
	HTTPOk      bool
	HTTPElapsed time.Duration
	HTTPErr     string
	HTTPSOk     bool
	HTTPSElapsed time.Duration
	HTTPSErr    string
}

// Addr 返回 "ip:port" 格式，用于表格与列表展示
func (r *ProxyResult) Addr() string {
	return r.IP + ":" + strconv.Itoa(r.Port)
}

// Proxy 从结果构造 HTTP 或 HTTPS 的 *Proxy（用于 Save）
func (r *ProxyResult) Proxy(protocol string) *Proxy {
	p := &Proxy{
		IP: r.IP, Port: r.Port, User: r.User, Passwd: r.Passwd, Protocol: protocol,
	}
	return p
}

var (
	dualResults   []*ProxyResult
	dualResultsMu sync.Mutex
)

// ClearDualResults 清空双协议结果列表（全量/轻量流程开始时调用）
func ClearDualResults() {
	dualResultsMu.Lock()
	dualResults = nil
	dualResultsMu.Unlock()
}

// AppendDualResult 追加一条双协议结果（由 collector 调用）
func AppendDualResult(r *ProxyResult) {
	dualResultsMu.Lock()
	dualResults = append(dualResults, r)
	dualResultsMu.Unlock()
}

// GetDualResults 返回当前双协议结果列表的副本，供 WriteTo 写 http+s 与表格
func GetDualResults() []*ProxyResult {
	dualResultsMu.Lock()
	out := make([]*ProxyResult, len(dualResults))
	copy(out, dualResults)
	dualResultsMu.Unlock()
	return out
}

// TruncateErr 将错误信息截断为最多 maxLen 字符（表格展示用）
func TruncateErr(s string, maxLen int) string {
	s = strings.TrimSpace(s)
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}
