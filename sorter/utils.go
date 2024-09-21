package sorter

import (
	"io"
	"io/fs"
	"os"
	"regexp"
)

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

func SortFiles(files []fs.DirEntry, conf Config) (err error) {
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		for _, r := range conf.Rules {
			if regexp.MustCompile(r.Pattern).MatchString(f.Name()) {
				_, err = moveFile(conf.Source+"/"+f.Name(), r.Target+"/"+f.Name())
				if err != nil {
					return
				}
			}
		}
	}
	return
}
