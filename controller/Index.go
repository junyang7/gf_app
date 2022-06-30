package controller

import (
	"github.com/junyang7/gf/engine/api"
	"github.com/junyang7/gf/engine/http"
	"github.com/junyang7/gf/engine/rpc"
	"github.com/junyang7/gf/tool/_json"
)

type Index struct {
}

func (this *Index) Rpc(ctx *rpc.Context) {

	res := map[string]interface{}{}
	res["debug"] = "rpc.Index.Rpc"
	ctx.Render(res)

}

func (this *Index) Api(ctx *api.Context) {

	res := map[string]interface{}{}
	res["debug"] = "api.Index.Api"
	ctx.Render(res)

}

func (this *Index) HttpApi(ctx *http.Context) {

	res := map[string]interface{}{}
	res["debug"] = "httpApi.Index.HttpApi"
	ctx.RenderJson(res)

}

func (this *Index) HttpByteList(ctx *http.Context) {

	res := map[string]interface{}{}
	res["debug"] = "HttpByteList"
	ctx.RenderByteList(_json.Encode(res))

}

func (this *Index) HttpCopy(ctx *http.Context) {

	res := "index.html2"
	ctx.RenderCopy(res)

}
