package main

import (
	"fmt"
	"git.ziji.fun/junyang7/gf/common"
	"git.ziji.fun/junyang7/gf/tool/_rpc"
)

func main() {

	addr := "127.0.0.1:10005"
	path := "/rpc/index"
	data := map[string]string{}
	res := common.NewResponse()
	_rpc.Request(addr, path, data, res)
	fmt.Println(res.Data)

}
