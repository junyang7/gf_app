package _file

func (this *File) Close() {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.closed {
		return
	}
	if err := this.file.Close(); nil != err {
		panic(err)
	}
	this.closed = true
}
