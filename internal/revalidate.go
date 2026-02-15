package internal

import (
	"bufio"
	"bytes"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

var revalidateProtos = []string{"http", "https", "socks4", "socks5"}

// RevalidateFromDir 从指定目录读取各协议列表，逐条复测（GET eastmoney + sinajs，2 秒内），通过则保留
func RevalidateFromDir(inputDir string) int {
	ClearDB()
	var total int
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
			if !CheckProxy(p) {
				continue
			}
			Save(p)
			total++
		}
	}
	return total
}
