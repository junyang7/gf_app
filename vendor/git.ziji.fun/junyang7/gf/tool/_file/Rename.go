package _file

import "os"

func Rename(source string, target string) {
	if err := os.Rename(source, target); nil != err {
		panic(err)
	}
}
