package _file

func (this *File) Write(content interface{}) {
	this.WriteOffset(content, 0)
}
