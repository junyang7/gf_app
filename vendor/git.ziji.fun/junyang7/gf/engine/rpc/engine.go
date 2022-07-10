package rpc

import (
	"context"
	"fmt"
	"git.ziji.fun/junyang7/gf/common/interceptor"
	"git.ziji.fun/junyang7/gf/engine/rpc/pb"
	"google.golang.org/grpc"
	"net"
)

type engine struct {
	conf *Conf
	pb.UnimplementedServiceServer
}

func Run(conf *Conf) {
	interceptor.
		Insure("" != conf.Network).
		Message("参数错误").
		Data(map[string]interface{}{"network": conf.Network}).
		Do()
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
	fmt.Println(conf)
	this := &engine{conf: conf}
	l, err := net.Listen(this.conf.Network, this.conf.Ip+":"+this.conf.Port)
	if nil != err {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterServiceServer(s, this)
	if err := s.Serve(l); nil != err {
		panic(err)
	}
}
func (this *engine) Call(c context.Context, r *pb.Request) (*pb.Response, error) {
	defer func() {
		if err := recover(); nil != err {
			fmt.Println(err)
		}
	}()
	ctx := newContext(r, this.conf)
	ctx.do()
	return ctx.render()
}
