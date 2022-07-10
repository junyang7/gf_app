package api

import (
	"git.ziji.fun/junyang7/gf/common"
	"git.ziji.fun/junyang7/gf/common/interceptor"
	"git.ziji.fun/junyang7/gf/tool/_json"
	"git.ziji.fun/junyang7/gf/tool/_slice"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"
)

type Context struct {
	w          http.ResponseWriter
	r          *http.Request
	timeS      time.Time
	timeE      time.Time
	response   *common.Response
	router     *router
	get        map[string]string
	Method     string
	Path       string
	Conf       *Conf
	HeaderList map[string]string
}

func newContext(w http.ResponseWriter, r *http.Request, conf *Conf) *Context {
	this := &Context{timeS: time.Now()}
	this.w = w
	this.r = r
	this.response = common.NewResponse()
	this.Conf = conf
	this.HeaderList = map[string]string{"content-type": "application/json"}
	this.Path = r.URL.Path
	this.Method = r.Method
	return this
}
func (this *Context) Render(data interface{}) {
	if nil != data {
		this.response.Data = data
	}
}
func (this *Context) Header(name string) string {
	return this.r.Header.Get(name)
}
func (this *Context) Get(name string) string {
	if value, ok := this.get[name]; ok {
		return value
	}
	return this.r.URL.Query().Get(name)
}
func (this *Context) Post(name string) string {
	return this.r.PostFormValue(name)
}
func (this *Context) Request(name string) string {
	if value := this.Post(name); value != "" {
		return value
	}
	return this.Get(name)
}
func (this *Context) Cookie(name string) string {
	cookie, err := this.r.Cookie(name)
	if nil != err {
		return ""
	}
	return cookie.Value
}
func (this *Context) do() {
	defer this.handleException()
	this.checkOrigin()
	this.parseRouter()
	interceptor.
		Insure(nil != this.router).
		Message("路由不存在").
		Data(map[string]string{"method": this.Method, "path": this.Path}).
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
func (this *Context) checkOrigin() {
	origin := this.Header("Origin")
	matchedList := regexp.MustCompile("(\\S+)://([^:]+):?(\\d+)?").FindStringSubmatch(strings.Trim(origin, "/"))
	if 0 == len(matchedList) {
		return
	}
	for _, origin := range this.Conf.Origin {
		if "*" == origin || matchedList[2] == origin || "." == origin[0:1] && matchedList[2][len(matchedList[2])-len(origin):] == origin {
			headerValue := matchedList[1] + "://" + matchedList[2]
			if 4 == len(matchedList) {
				headerValue += ":" + matchedList[3]
			}
			this.HeaderList["access-control-allow-origin"] = headerValue
			return
		}
	}
	interceptor.Insure(false).
		Message("跨域阻止").
		Data(map[string]interface{}{"origin": origin}).
		Do()
}
func (this *Context) parseRouter() {
	for _, r := range routerList {
		uri := r.prefix + r.path
		if 0 == len(r.regexp) {
			if uri == this.Path && (0 == len(r.methodList) || _slice.In(this.Method, routerMethodList)) {
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
				if 0 == len(r.methodList) || _slice.In(this.Method, routerMethodList) {
					for index, name := range nameList {
						this.get[name] = matchedList[index+1]
					}
					this.router = r
				}
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
func (this *Context) render() {
	for k, v := range this.HeaderList {
		this.w.Header().Set(k, v)
	}
	this.time()
	this.w.Write(_json.Encode(this.response))
}
