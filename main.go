package main

import (
	"gf_app/conf"
	_ "gf_app/router"
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

	wg.Wait()

}
