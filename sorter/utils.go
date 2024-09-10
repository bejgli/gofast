package sorter

import (
	"io"
	"os"
)

func MoveFile(source string, target string) (written int64, err error) {
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
