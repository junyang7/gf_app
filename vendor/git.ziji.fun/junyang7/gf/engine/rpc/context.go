package rpc

import (
	"git.ziji.fun/junyang7/gf/common"
	"git.ziji.fun/junyang7/gf/common/interceptor"
	"git.ziji.fun/junyang7/gf/engine/rpc/pb"
	"git.ziji.fun/junyang7/gf/tool/_json"
	"reflect"
	"regexp"
	"strings"
	"time"
)

type Context struct {
	r        *pb.Request
	timeS    time.Time
	timeE    time.Time
	response *common.Response
	router   *router
	get      map[string]string
	Path     string
	Conf     *Conf
}

func newContext(r *pb.Request, conf *Conf) *Context {
	this := &Context{timeS: time.Now()}
	this.r = r
	this.response = common.NewResponse()
	this.Conf = conf
	return this
}
func (this *Context) Render(data interface{}) {
	if nil != data {
		this.response.Data = data
	}
}
func (this *Context) Request(name string) string {
	if v, ok := this.r.Data[name]; ok {
		return v
	}
	return ""
}
func (this *Context) do() {
	defer this.handleException()
	this.Path = this.r.Path
	this.parseRouter()
	interceptor.
		Insure(nil != this.router).
		Message("路由不存在").
		Data(map[string]string{"path": this.Path}).
		Do()
	this.middlewareBefore()
	this.business()
	this.middlewareAfter()
}
func (this *Context) handleException() {
	if err := recover(); nil != err {
		switch err.(type) {
		case *common.Response:
			response := err.(*common.Response)
			this.response.Code = response.Code
			this.response.Message = response.Message
			this.response.Data = response.Data
		default:
			this.response.Code = common.FailureCode
			this.response.Message = common.FailureMessage
		}
	}
}
func (this *Context) parseRouter() {
	for _, r := range routerList {
		uri := r.prefix + r.path
		if 0 == len(r.regexp) {
			if uri == this.Path {
				this.router = r
				break
			}
		} else {
			nameList := make([]string, 0, len(r.regexp))
			for name, pattern := range r.regexp {
				if nameRegexp := "{" + name + "}"; strings.Index(uri, nameRegexp) > 0 {
					uri = strings.Replace(uri, nameRegexp, pattern, -1)
					nameList = append(nameList, name)
				}
			}
			if matchedList := regexp.MustCompile(uri).FindStringSubmatch(this.Path); len(matchedList) > 0 {
				for index, name := range nameList {
					this.get[name] = matchedList[index+1]
				}
				this.router = r
			}
		}
	}
}
func (this *Context) middlewareBefore() {
	for _, middleware := range this.router.middlewareList {
		if middleware.before.IsValid() {
			var parameter []reflect.Value
			if middleware.before.Type().NumIn() > 0 {
				parameter = append(parameter, reflect.ValueOf(this))
			}
			middleware.before.Call(parameter)
		}
	}
}
func (this *Context) business() {
	var parameter []reflect.Value
	if this.router.action.Type().NumIn() > 0 {
		parameter = append(parameter, reflect.ValueOf(this))
	}
	this.router.action.Call(parameter)
}
func (this *Context) middlewareAfter() {
	for _, middleware := range this.router.middlewareList {
		if middleware.after.IsValid() {
			var parameter []reflect.Value
			if middleware.after.Type().NumIn() > 0 {
				parameter = append(parameter, reflect.ValueOf(this))
			}
			middleware.after.Call(parameter)
		}
	}
}
func (this *Context) time() {
	this.timeE = time.Now()
	this.response.Time = this.timeE.Unix()
	this.response.Consume = this.timeE.Sub(this.timeS).Microseconds()
}
func (this *Context) render() (*pb.Response, error) {
	this.time()
	res := &pb.Response{Response: _json.Encode(this.response)}
	return res, nil
}
