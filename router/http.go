package router

import (
	"gf_app/controller"
	"git.ziji.fun/junyang7/gf/engine/http"
)

func init() {
	http.Router.Any("/http/api", &controller.Index{}, "HttpApi")
	http.Router.Any("/http/byteList", &controller.Index{}, "HttpByteList")
	http.Router.Any("/http/copy", &controller.Index{}, "HttpCopy")
}
