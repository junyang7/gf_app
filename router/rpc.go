package router

import (
	"gf_app/controller"
	"github.com/junyang7/gf/engine/rpc"
)

func init() {
	rpc.Router.Post("/rpc/index", &controller.Index{}, "Index")
}
