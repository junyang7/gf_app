package http

import (
	"fmt"
	"git.ziji.fun/junyang7/gf/common/interceptor"
	"net/http"
)

type engine struct {
	conf *Conf
}

func Run(conf *Conf) {
	interceptor.
		Insure("" != conf.Ip).
		Message("参数错误").
		Data(map[string]interface{}{"ip": conf.Ip}).
		Do()
	interceptor.
		Insure("" != conf.Port).
		Message("参数错误").
		Data(map[string]interface{}{"port": conf.Port}).
		Do()
	interceptor.
		Insure("" != conf.Root).
		Message("参数错误").
		Data(map[string]interface{}{"root": conf.Root}).
		Do()
	fmt.Println(conf)
	this := &engine{conf: conf}
	addr := this.conf.Ip + ":" + this.conf.Port
	err := http.ListenAndServe(addr, this)
	if nil != err {
		panic(err)
	}
}
func (this *engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "OPTIONS" == r.Method {
		return
	}
	defer func() {
		if err := recover(); nil != err {
			fmt.Println(err)
		}
	}()
	ctx := newContext(w, r, this.conf)
	ctx.do()
	ctx.render()
}
