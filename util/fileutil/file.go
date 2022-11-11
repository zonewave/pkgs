package fileutil

import (
	"os"
	"path/filepath"
	"strings"
)

func AbsPath(inPath string) string {
	if strings.HasPrefix(inPath, "$") {
		end := strings.Index(inPath, string(os.PathSeparator))
		inPath = os.Getenv(inPath[1:end]) + inPath[end:]
	}

	if filepath.IsAbs(inPath) {
		return filepath.Clean(inPath)
	}

	p, err := filepath.Abs(inPath)
	if err == nil {
		return filepath.Clean(p)
	}

	return ""
}

func FileInfo(path string) (name, ext string) {
	if fileExt := filepath.Ext(path); len(fileExt) > 1 {
		ext = fileExt[1:]
	}

	name = strings.TrimSuffix(path, filepath.Ext(path))
	return
}
