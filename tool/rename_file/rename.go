package main

import (
	"github.com/cockroachdb/errors"
	"github.com/zonewave/pkgs/log"
	"github.com/zonewave/pkgs/mstrings"
	"github.com/zonewave/pkgs/slice"
	"os"
	"path/filepath"
)

func main() {
	run(os.Args)
}

func run(args []string) {
	if len(args) < 2 {
		log.Fatal("param less than 1 arguments")
		return
	}

	err := renameFiles(os.Args[1])
	if err != nil {
		log.Error(err.Error())
		return
	}

	return
}

func renameFiles(curPath string) error {
	dirEntry, err := os.ReadDir(curPath)
	if err != nil {
		return errors.WithStack(err)
	}
	slice.Slices[os.DirEntry](dirEntry).IterFn(
		func(dirEntryNode os.DirEntry) bool {
			if dirEntryNode.IsDir() || filepath.Ext(dirEntryNode.Name()) != ".py" {
				return true
			}
			fileName := dirEntryNode.Name()
			fileNameNoExt := fileName[:len(fileName)-3]
			start := -1
			for i := len(fileNameNoExt) - 1; i >= 0; i-- {
				if mstrings.CharIsDigital(fileNameNoExt[i]) {
					continue
				} else {
					start = i
					break
				}
			}
			if start == -1 || start == len(fileNameNoExt)-1 {
				return true
			}
			pNum := fileNameNoExt[start+1:]
			before := fileNameNoExt[:start+1]
			err = os.Rename(fileName, "_"+pNum+"_"+before+".py")
			if err != nil {
				err = errors.WithStack(err)
				return false
			}
			return true
		})

	if err != nil {
		return errors.WithStack(err)
	}

	return nil

}
