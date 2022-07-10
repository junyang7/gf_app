package _file

import "os"

func Exists(filepath string) bool {
	f, err := os.Stat(filepath)
	if nil != err {
		return false
	}
	return !f.IsDir()
}
