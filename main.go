package main

import (
	"gf_app/conf"
	_ "gf_app/router"
	"github.com/junyang7/gf/engine/api"
	"github.com/junyang7/gf/engine/http"
	"github.com/junyang7/gf/engine/rpc"
	"github.com/junyang7/gf/engine/web"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		web.Run(conf.Application.Web)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		rpc.Run(conf.Application.Rpc)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		api.Run(conf.Application.Api)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		http.Run(conf.Application.Http)
		wg.Done()
	}()

	wg.Wait()

}
