package _as

func ByteList(data interface{}) []byte {
	return []byte(String(data))
}
