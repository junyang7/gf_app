package _file

import (
	"io/ioutil"
	"os"
)

func (this *File) Read() []byte {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.file.Seek(0, 0)
	b, err := ioutil.ReadAll(this.file)
	if nil != err {
		panic(err)
	}
	return b
}

func Read(filepath string) []byte {
	b, err := os.ReadFile(filepath)
	if nil != err {
		panic(err)
	}
	return b
}
