package _file

import (
	"os"
	"sync"
)

type File struct {
	file   *os.File
	lock   *sync.Mutex
	closed bool
}
