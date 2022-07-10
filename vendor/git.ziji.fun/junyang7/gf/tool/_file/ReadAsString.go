package _file

func (this *File) ReadAsString() string {
	return string(this.Read())
}

func ReadAsString(filepath string) string {
	return string(Read(filepath))
}
