package _json

import "encoding/json"

func Encode(data interface{}) []byte {
	b, err := json.Marshal(data)
	if nil != err {
		panic(err)
	}
	return b
}
