package _json

func EncodeAsString(data interface{}) string {
	return string(Encode(data))
}
