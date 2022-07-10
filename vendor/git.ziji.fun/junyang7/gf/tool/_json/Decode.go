package _json

import "encoding/json"

func Decode(source []byte, target interface{}) {
	if err := json.Unmarshal(source, target); nil != err {
		panic(err)
	}
}
