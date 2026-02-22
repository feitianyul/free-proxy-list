package internal

import (
	"bufio"
	"bytes"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

var revalidateProtos = []string{"http", "https"}

// RevalidateFromDir 从指定目录读取各协议列表，双协议并发复测（五域名中任意 3 个在 2 秒内成功则保留）
func RevalidateFromDir(inputDir string) int {
	ClearDB()
	ClearDualResults()
	var candidates []*Proxy
	for _, proto := range revalidateProtos {
		path := filepath.Join(inputDir, proto+".txt")
		buf, err := os.ReadFile(path)
		if err != nil {
			slog.Debug("revalidate: skip file", "path", path, "err", err)
			continue
		}
		s := bufio.NewScanner(bytes.NewReader(buf))
		for s.Scan() {
			line := strings.TrimSpace(s.Text())
			if line == "" {
				continue
			}
			p, err := ParseProxyURL(proto, line)
			if err != nil {
				continue
			}
			candidates = append(candidates, p)
		}
	}
	return ValidateProxiesDual(candidates, GetCheckWorkers())
}
