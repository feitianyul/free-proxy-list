package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const debugLogPath = ".cursor/debug.log"

var debugLogMu sync.Mutex

// DebugLog writes one NDJSON line to .cursor/debug.log for hypothesis tracking.
func DebugLog(location, message string, data map[string]interface{}, hypothesisId string) {
	debugLogMu.Lock()
	defer debugLogMu.Unlock()
	dir := filepath.Dir(debugLogPath)
	_ = os.MkdirAll(dir, 0755)
	f, err := os.OpenFile(debugLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	payload := map[string]interface{}{
		"id":          "log_" + time.Now().Format("20060102150405"),
		"timestamp":   time.Now().UnixMilli(),
		"location":    location,
		"message":     message,
		"data":        data,
		"hypothesisId": hypothesisId,
	}
	b, _ := json.Marshal(payload)
	f.Write(b)
	f.Write([]byte("\n"))
}
