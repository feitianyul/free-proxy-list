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

// 只保留 http、https、socks4、socks5 三种代理
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

	// Generate total.svg and update README.md
	WriteTotalAndUpdateReadme(dir, counters)
}

func WriteTotalAndUpdateReadme(dir string, counters map[string]int) {
	// Calculate total
	total := 0
	protocols := make([]string, 0, len(counters))
	for proto, count := range counters {
		protocols = append(protocols, proto)
		total += count
	}
	sort.Strings(protocols)

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

	// Build table content for README replacement
	var tableContent strings.Builder
	for _, proto := range protocols {
		count := counters[proto]
		url := fmt.Sprintf("https://raw.githubusercontent.com/wiki/gfpcom/free-proxy-list/lists/%s.txt", proto)
		tableContent.WriteString(fmt.Sprintf("| %s | %d | %s |\n",
			strings.ToUpper(proto),
			count,
			url))
	}

	startMarker := "<!-- BEGIN PROXY LIST -->"
	endMarker := "<!-- END PROXY LIST -->"

	// Update README.md (English)
	readmePath := filepath.Join(dir, "..", "README.md")
	readmeContent, err := os.ReadFile(readmePath)
	if err == nil {
		newSection := fmt.Sprintf(`
Last Updated: %s (%s)

**Total Proxies: %d**

Click on your preferred proxy type to get the latest list. These links always point to the most recently updated proxy files.

| Protocol | Count | Download |
|----------|-------|----------|
%s`, tsUTC, tsUTC8, total, tableContent.String())
		replaceReadmeSection(readmePath, string(readmeContent), startMarker, endMarker, newSection)
	}

	// Update README_ZH.md (中文)
	readmeZHPath := filepath.Join(dir, "..", "README_ZH.md")
	readmeZHContent, errZH := os.ReadFile(readmeZHPath)
	if errZH == nil {
		newSectionZH := fmt.Sprintf(`
最后更新：%s（%s）

**代理总数：%d**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
%s`, tsUTC, tsUTC8, total, tableContent.String())
		replaceReadmeSection(readmeZHPath, string(readmeZHContent), startMarker, endMarker, newSectionZH)
	}
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
