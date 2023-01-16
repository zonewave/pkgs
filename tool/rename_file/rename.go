package main

import (
	"github.com/cockroachdb/errors"
	"github.com/zonewave/pkgs/log"
	"github.com/zonewave/pkgs/mstrings"
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
	for _, dirEntryNode := range dirEntry {
		if dirEntryNode.IsDir() || filepath.Ext(dirEntryNode.Name()) != ".py" {
			continue
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
			continue
		}
		pNum := fileNameNoExt[start+1:]
		before := fileNameNoExt[:start+1]
		err = os.Rename(fileName, "_"+pNum+"_"+before+".py")
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil

}
