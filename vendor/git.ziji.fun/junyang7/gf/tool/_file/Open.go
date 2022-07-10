package _file

import (
	"os"
	"sync"
)

// Open 打开文件
func Open(filepath string, flag int, perm os.FileMode) *File {
	this := &File{
		lock: &sync.Mutex{},
	}
	file, err := os.OpenFile(filepath, flag, perm)
	if nil != err {
		panic(err)
	}
	this.file = file
	return this
}
