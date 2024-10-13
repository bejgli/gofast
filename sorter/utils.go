package sorter

import (
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func SortFiles(files []fs.DirEntry, conf Config) (err error) {
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		for _, r := range conf.Rules {
			dstPath := filepath.Join(r.Target, f.Name())
			if checkExisting(dstPath) && !r.Overwrite {
				log.Println("Skipping file: ", f.Name())
				continue
			}

			srcPath := filepath.Join(conf.Source, f.Name())
			if regexp.MustCompile(r.Pattern).MatchString(f.Name()) {
				_, err = moveFile(srcPath, dstPath)
				if err != nil {
					return
				}
			}
		}
	}
	return
}

func checkExisting(dstPath string) bool {
	_, err := os.Stat(dstPath)
	if err != nil && errors.Is(err, fs.ErrNotExist) {
		return false
	}
	return true
}

func moveFile(source string, target string) (written int64, err error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return
	}
	defer sourceFile.Close()

	targetFile, err := os.Create(target)
	if err != nil {
		return
	}
	defer targetFile.Close()

	written, err = io.Copy(targetFile, sourceFile)
	if err != nil {
		return
	}

	err = os.Remove(source)
	if err != nil {
		return
	}

	return
}
