package router

import (
	"gf_app/controller"
	"git.ziji.fun/junyang7/gf/engine/rpc"
)

func init() {
	rpc.Router.Any("/rpc/index", &controller.Index{}, "Rpc")
}
