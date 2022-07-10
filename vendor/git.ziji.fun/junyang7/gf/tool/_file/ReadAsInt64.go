package _file

import "git.ziji.fun/junyang7/gf/tool/_as"

func (this *File) ReadAsInt64() int64 {
	return _as.Int64(this.Read())
}

func ReadAsInt64(filepath string) int64 {
	return _as.Int64(Read(filepath))
}
