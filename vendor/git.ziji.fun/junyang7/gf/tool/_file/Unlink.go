package _file

import "os"

func Unlink(filepath string) {
	_ = os.Remove(filepath)
}
