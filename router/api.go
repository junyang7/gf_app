package router

import (
	"gf_app/controller"
	"git.ziji.fun/junyang7/gf/engine/api"
)

func init() {
	api.Router.Any("/api/index", &controller.Index{}, "Api")
}
