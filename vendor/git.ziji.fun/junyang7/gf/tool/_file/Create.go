package _file

import "os"

func Create(filepath string) {
	_, err := os.Create(filepath)
	if nil != err {
		panic(err)
	}
}
