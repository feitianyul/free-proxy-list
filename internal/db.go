package internal

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var (
	db = make(map[string]*Proxy)
)

// ClearDB 清空内存中的代理库，用于轻量复测前重置
func ClearDB() {
	db = make(map[string]*Proxy)
}

// 只保留 http、https 两种代理
func Save(it *Proxy) {
	if !IsAllowedProtocol(it.Protocol) {
		return
	}
	h := md5.New()
	id := hex.EncodeToString(h.Sum([]byte(fmt.Sprintf("%s://%s:%v", it.Protocol, it.IP, it.Port))))
	db[id] = it
}

func WriteTo(dir string) {
	files := make(map[string]*os.File)
	defer func() {
		for _, f := range files {
			f.Sync() // nolint: errcheck
			f.Close()
		}
	}()

	// Get all keys and sort them
	keys := make([]string, 0, len(db))
	for k := range db {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	counters := make(map[string]int)

	// Iterate through sorted keys
	for _, key := range keys {
		it := db[key]
		file, ok := files[it.Protocol]
		if !ok {
			file, _ = os.Create(filepath.Join(dir, it.Protocol+".txt"))
			files[it.Protocol] = file
		}

		c, ok := counters[it.Protocol]
		if !ok {
			counters[it.Protocol] = 1
		} else {
			counters[it.Protocol] = c + 1
		}

		file.WriteString(it.String() + "\n") // nolint: errcheck
	}

	// 保证 http.txt / https.txt 存在（即使数量为 0，便于 README 下载链接一致）
	for _, proto := range []string{"http", "https"} {
		if _, ok := files[proto]; !ok {
			if f, err := os.Create(filepath.Join(dir, proto+".txt")); err == nil {
				files[proto] = f
			}
		}
	}

	results := GetDualResults()
	// 通过测试：与表格同序，先构建 passedResults 再按 IP:port 去重，再写 passed.txt
	passedResults := make([]*ProxyResult, 0, len(results))
	for _, r := range results {
		if r.HTTPOk || r.HTTPSOk {
			passedResults = append(passedResults, r)
		}
	}
	// 按代理地址去重，保留首次出现（与表格、passed.txt 一致）
	seen := make(map[string]bool)
	passedResultsDedup := make([]*ProxyResult, 0, len(passedResults))
	for _, r := range passedResults {
		addr := r.Addr()
		if seen[addr] {
			continue
		}
		seen[addr] = true
		passedResultsDedup = append(passedResultsDedup, r)
	}
	passedFileName := "passed.txt"
	if f, err := os.Create(filepath.Join(dir, passedFileName)); err == nil {
		for _, r := range passedResultsDedup {
			f.WriteString("http://" + r.Addr() + "\n") // nolint: errcheck
		}
		f.Sync()  // nolint: errcheck
		f.Close()
	}
	counters["passed"] = len(passedResultsDedup)

	// Generate total.svg and update README (list section + proxy table)
	WriteTotalAndUpdateReadme(dir, counters, results, passedResultsDedup)
}

func WriteTotalAndUpdateReadme(dir string, counters map[string]int, results []*ProxyResult, passedResults []*ProxyResult) {
	// 固定展示两种协议 + 通过测试（HTTP / HTTPS / 通过测试）
	protocolOrder := []string{"http", "https"}
	for _, p := range protocolOrder {
		if _, ok := counters[p]; !ok {
			counters[p] = 0
		}
	}
	total := 0
	for _, p := range protocolOrder {
		total += counters[p]
	}
	if _, ok := counters["passed"]; !ok {
		counters["passed"] = 0
	}

	tsUTC := time.Now().UTC().Format("2006-01-02 15:04:05 UTC")
	tsUTC8 := time.Now().In(time.FixedZone("UTC+8", 8*3600)).Format("2006-01-02 15:04:05 UTC+8")

	// Generate total.svg using shields.io
	svgURL := fmt.Sprintf("https://img.shields.io/badge/total-%d-blue", total)
	resp, err := httpGet(svgURL)
	if err == nil && resp != nil {
		defer resp.Close()
		outPath := filepath.Join(dir, "total.svg")
		_ = os.WriteFile(outPath, resp.Bytes(), 0644) // nolint: errcheck
	}

	// Build protocol table for README (协议 | 数量 | 下载)：HTTP / HTTPS / 通过测试
	var tableContent strings.Builder
	for _, proto := range protocolOrder {
		count := counters[proto]
		fileName := proto + ".txt"
		url := fmt.Sprintf("https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/%s", fileName)
		tableContent.WriteString(fmt.Sprintf("| %s | %d | %s |\n",
			strings.ToUpper(proto),
			count,
			url))
	}
	tableContent.WriteString(fmt.Sprintf("| 通过测试 (Passed) | %d | %s |\n",
		counters["passed"],
		"https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt"))

	startMarker := "<!-- BEGIN PROXY LIST -->"
	endMarker := "<!-- END PROXY LIST -->"

	// Update README.md（中文，GitHub 首页默认显示）
	readmePath := filepath.Join(dir, "..", "README.md")
	readmeContent, err := os.ReadFile(readmePath)
	if err == nil {
		newSectionZH := fmt.Sprintf(`
最后更新：%s（%s）

**代理总数：%d**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
%s`, tsUTC, tsUTC8, total, tableContent.String())
		replaceReadmeSection(readmePath, string(readmeContent), startMarker, endMarker, newSectionZH)
		readmeContent, _ = os.ReadFile(readmePath)
		writeReadmeProxyTable(readmePath, string(readmeContent), passedResults, true)
	}

	// Update README_EN.md (English)
	readmeENPath := filepath.Join(dir, "..", "README_EN.md")
	readmeENContent, errEN := os.ReadFile(readmeENPath)
	if errEN == nil {
		newSection := fmt.Sprintf(`
Last Updated: %s (%s)

**Total Proxies: %d**

Click on your preferred proxy type to get the latest list. These links always point to the most recently updated proxy files.

| Protocol | Count | Download |
|----------|-------|----------|
%s`, tsUTC, tsUTC8, total, tableContent.String())
		replaceReadmeSection(readmeENPath, string(readmeENContent), startMarker, endMarker, newSection)
		readmeENContent, _ = os.ReadFile(readmeENPath)
		writeReadmeProxyTable(readmeENPath, string(readmeENContent), passedResults, false)
	}
}

const tableErrMaxLen = 20

func writeReadmeProxyTable(readmePath, content string, results []*ProxyResult, zh bool) {
	startMarker := "<!-- BEGIN PROXY TABLE -->"
	endMarker := "<!-- END PROXY TABLE -->"
	startIdx := strings.Index(content, startMarker)
	endIdx := strings.Index(content, endMarker)
	if startIdx == -1 || endIdx == -1 {
		return
	}
	maxRows := 100
	if len(results) < maxRows {
		maxRows = len(results)
	}
	var table strings.Builder
	if zh {
		table.WriteString("| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |\n")
		table.WriteString("|----------|---------------|------------|----------------------|-------------------|---------------------|------|\n")
	} else {
		table.WriteString("| Address | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | Protocol |\n")
		table.WriteString("|---------|---------------|------------|----------------------|-------------------|---------------------|----------|\n")
	}
	for i := 0; i < maxRows; i++ {
		r := results[i]
		elapsed := r.HTTPElapsed
		if r.Protocol == "https" && len(r.HTTPSElapsed) >= 5 {
			elapsed = r.HTTPSElapsed
		} else if (r.Protocol == "http" || r.Protocol == "http/s") && len(r.HTTPElapsed) >= 5 {
			elapsed = r.HTTPElapsed
		}
		c1 := formatCellElapsed(elapsed, 0)
		c2 := formatCellElapsed(elapsed, 1)
		c3 := formatCellElapsed(elapsed, 2)
		c4 := formatCellElapsed(elapsed, 3)
		c5 := formatCellElapsed(elapsed, 4)
		table.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n", r.Addr(), c1, c2, c3, c4, c5, r.Protocol))
	}
	before := content[:startIdx+len(startMarker)]
	after := content[endIdx:]
	newContent := before + "\n" + table.String() + "\n" + after
	_ = os.WriteFile(readmePath, []byte(newContent), 0644) // nolint: errcheck
}

const tableCellMaxMs = 2000 // 与 check.checkTimeout 一致，≤此值且>0 显示 ✓Xms，否则显示 否

func formatCellElapsed(elapsed []time.Duration, i int) string {
	if i >= len(elapsed) {
		return "否"
	}
	d := elapsed[i]
	if d > 0 && d.Milliseconds() <= tableCellMaxMs {
		return fmt.Sprintf("✓ %dms", d.Milliseconds())
	}
	return "否"
}

func formatCell(ok bool, elapsed time.Duration, errMsg string) string {
	if ok {
		return fmt.Sprintf("✓ %dms", elapsed.Milliseconds())
	}
	return "否 " + TruncateErr(errMsg, tableErrMaxLen)
}

func replaceReadmeSection(readmePath, content, startMarker, endMarker, newSection string) {
	startIdx := strings.Index(content, startMarker)
	endIdx := strings.Index(content, endMarker)
	if startIdx != -1 && endIdx != -1 {
		before := content[:startIdx+len(startMarker)]
		after := content[endIdx:]
		newContent := before + "\n" + newSection + "\n" + after
		_ = os.WriteFile(readmePath, []byte(newContent), 0644) // nolint: errcheck
	}
}

// httpGet fetches URL and returns a small wrapper with the body bytes and Close()
type respWrap struct {
	b []byte
}

func (r *respWrap) Close() error  { return nil }
func (r *respWrap) Bytes() []byte { return r.b }

func httpGet(url string) (*respWrap, error) {
	// Use simple http.Get but avoid adding net/http import at top if not present; import locally
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &respWrap{b: data}, nil
}
