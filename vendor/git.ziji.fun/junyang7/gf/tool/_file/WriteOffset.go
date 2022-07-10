package _file

import "git.ziji.fun/junyang7/gf/tool/_as"

func (this *File) WriteOffset(content interface{}, offset int64) {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, err := this.file.WriteAt([]byte(_as.String(content)), offset)
	if err != nil {
		panic(err)
	}
}
