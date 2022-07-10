package _file

import (
	"io"
	"os"
)

func Copy(sourceFilepath string, targetFilePath string) {
	sourceFile, err := os.OpenFile(sourceFilepath, os.O_RDONLY, os.ModePerm)
	if nil != err {
		panic(err)
	}
	defer sourceFile.Close()
	targetFile, err := os.OpenFile(targetFilePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if nil != err {
		panic(err)
	}
	defer targetFile.Close()
	if _, err := io.Copy(targetFile, sourceFile); nil != err {
		panic(err)
	}
}
