package controller

import "github.com/junyang7/gf/engine/rpc"

type Index struct {
}

func (this *Index) Index(ctx *rpc.Context) {

	res := map[string]interface{}{}
	res["debug"] = "rpc.Index.Index"
	ctx.Render(res)

}
