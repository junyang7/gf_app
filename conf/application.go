package conf

import (
	"github.com/junyang7/gf/engine/http"
	"github.com/junyang7/gf/engine/rpc"
	"github.com/junyang7/gf/engine/web"
	"github.com/junyang7/gf/tool/_file"
	"github.com/junyang7/gf/tool/_json"
)

var Application *struct {
	Http *http.Conf `json:"http"`
	Web  *web.Conf  `json:"web"`
	Rpc  *rpc.Conf  `json:"rpc"`
}

func init() {
	filepath := "application.json"
	conf := []byte("{}")
	if _file.Exists(filepath) {
		conf = _file.Read(filepath)
	}
	_json.Decode(conf, &Application)
}
