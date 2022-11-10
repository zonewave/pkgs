package fileutil

import (
	"github.com/cockroachdb/errors"
	"github.com/spf13/afero"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	// ErrNotFound file not found
	ErrNotFound = errors.New("file not found")
)

// Afero  file interface
type Afero interface {
	afero.Fs
	ReadDir(dirname string) ([]os.FileInfo, error)
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, data []byte, perm os.FileMode) error
	WriteReader(path string, r io.Reader) (err error)
	Exists(path string) (bool, error)
}

// AbsPath  Absolute path
func AbsPath(inPath string) (string, error) {
	if strings.HasPrefix(inPath, "$") {
		end := strings.Index(inPath, string(os.PathSeparator))
		inPath = os.Getenv(inPath[1:end]) + inPath[end:]
	}
	return filepath.Abs(inPath)
}

// FileNameAndExt return file and fileExt
func FileNameAndExt(fileName string) (name, ext string) {
	for i := len(fileName) - 1; i >= 0 && !os.IsPathSeparator(fileName[i]); i-- {
		if fileName[i] == '.' {
			name = fileName[:i]
			ext = fileName[i+1:]
			break
		}
	}
	return

}

// FileExtNoDot return file extension without dot
func FileExtNoDot(fileName string) (ext string) {
	ext = filepath.Ext(fileName)
	if len(ext) > 0 {
		return ext[1:]
	}
	return
}

// SearchInPaths search file in paths
func SearchInPaths(fs Afero, paths []string, baseFileName string) (string, error) {
	for _, path := range paths {
		fileName := filepath.Join(path, baseFileName)
		if b, _ := fs.Exists(fileName); b {
			return fileName, nil
		}
	}
	return "", errors.WithStack(ErrNotFound)
}
