package conf

import (
	"git.ziji.fun/junyang7/gf/engine/api"
	"git.ziji.fun/junyang7/gf/engine/http"
	"git.ziji.fun/junyang7/gf/engine/rpc"
	"git.ziji.fun/junyang7/gf/engine/web"
	"git.ziji.fun/junyang7/gf/tool/_file"
	"git.ziji.fun/junyang7/gf/tool/_json"
)

var Application *struct {
	Http *http.Conf `json:"http"`
	Web  *web.Conf  `json:"web"`
	Rpc  *rpc.Conf  `json:"rpc"`
	Api  *api.Conf  `json:"api"`
}

func init() {
	filepath := "application.json"
	conf := []byte("{}")
	if _file.Exists(filepath) {
		conf = _file.Read(filepath)
	}
	_json.Decode(conf, &Application)
}
