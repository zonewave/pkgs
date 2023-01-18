package main

import (
	"github.com/cockroachdb/errors"
	"github.com/zonewave/pkgs/log"
	"github.com/zonewave/pkgs/mstrings"
	"github.com/zonewave/pkgs/slice"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	run(os.Args)
}

var renameFns = []DirEntryNodeFn{
	RenameMethod1, RenameMethod2,
}

func run(args []string) {
	if len(args) < 2 {
		log.Fatal("param less than 1 arguments")
		return
	}
	fn := RenameMethod1
	if len(args) == 3 {
		index, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatal("param fn index failed")
			return
		}
		fn = renameFns[index]
	}
	log.Infof("run start,param:%s", args[1:])
	err := renameFiles(args[1], fn)
	if err != nil {
		log.Error(err.Error())
		return
	}
	log.Infof("run end,param:%s", args[1:])
	return
}

func renameFiles(curPath string, fn DirEntryNodeFn) error {
	dirEntry, err := os.ReadDir(curPath)
	if err != nil {
		return errors.WithStack(err)
	}
	slice.Slices[os.DirEntry](dirEntry).IterFn(
		func(dirEntryNode os.DirEntry) bool {
			if dirEntryNode.IsDir() || filepath.Ext(dirEntryNode.Name()) != ".py" {
				return true
			}
			err = fn(dirEntryNode)
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

type DirEntryNodeFn func(dirEntryNode os.DirEntry) error

func RenameMethod1(dirEntryNode os.DirEntry) error {
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
		return nil
	}
	pNum := fileNameNoExt[start+1:]
	before := fileNameNoExt[:start+1]
	err := os.Rename(fileName, "_"+pNum+"_"+before+".py")
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
func RenameMethod2(dirEntryNode os.DirEntry) error {
	fileName := dirEntryNode.Name()
	newName := strings.Replace(fileName, "剑指", "", 1)
	newName = mstrings.SpaceRemoveAll(newName)
	err := os.Rename(fileName, newName)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
