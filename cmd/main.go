package main

import (
	"flag"
	"io/fs"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/gfpcom/free-proxy-list/internal"
)

var (
	dir         string
	revalidate  bool
	inputDir    string
	checkWorkers int
)

func main() {
	flag.StringVar(&dir, "dir", ".", "work directory (output list dir)")
	flag.BoolVar(&revalidate, "revalidate", false, "lightweight mode: read existing lists from -input-dir, re-check each proxy, write to -dir")
	flag.StringVar(&inputDir, "input-dir", "", "input directory for -revalidate (e.g. ../wiki/lists)")
	flag.IntVar(&checkWorkers, "check-workers", 0, "concurrent proxy check workers (0=default 2000, max 4000); env GFP_CHECK_WORKERS overrides default")
	flag.Parse()

	if checkWorkers != 0 {
		internal.CheckWorkers = checkWorkers
	}

	os.MkdirAll(filepath.Join(dir, "list"), 0755) // nolint: errcheck

	if revalidate {
		if inputDir == "" {
			inputDir = filepath.Join(dir, "list")
		}
		log.Println("revalidate from", inputDir, "->", dir)
		n := internal.RevalidateFromDir(inputDir)
		log.Println("revalidate passed:", n)
		internal.WriteTo(filepath.Join(dir, "list"))
		return
	}

	internal.ClearDualResults()

	// 只处理 http、https 两种代理源
	allowedSources := map[string]bool{"http": true, "https": true}

	err := fs.WalkDir(os.DirFS(filepath.Join(dir, "sources")), ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			slog.Warn("gfp: open source", slog.String("file", path), slog.Any("err", err))
			return nil
		}

		if d.IsDir() {
			return nil
		}

		// Get filename without extension
		filename := d.Name()
		proto := strings.ToLower(strings.TrimSuffix(filename, filepath.Ext(filename)))
		if !allowedSources[proto] {
			return nil
		}

		buf, err := os.ReadFile(filepath.Join(dir, "sources", path))
		if err != nil {
			slog.Warn("gfp: read source", slog.String("file", path), slog.Any("err", err))
			return nil
		}

		log.Println("--------" + path + "-------")
		err = internal.Load(proto, buf)
		if err != nil {
			slog.Warn("gfp: read source", slog.String("file", path), slog.Any("err", err))
			return nil
		}
		log.Println("---------------------------")
		log.Println("")

		return nil
	})

	internal.WriteTo(filepath.Join(dir, "list"))

	if err != nil {
		panic(err)
	}
}
