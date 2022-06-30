package router

import (
	"gf_app/controller"
	"github.com/junyang7/gf/engine/api"
)

func init() {
	api.Router.Any("/api/index", &controller.Index{}, "Api")
}
