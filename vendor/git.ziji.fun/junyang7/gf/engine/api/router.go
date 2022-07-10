package api

import (
	"git.ziji.fun/junyang7/gf/common/interceptor"
	"git.ziji.fun/junyang7/gf/tool/_slice"
	"reflect"
)

type middleware struct {
	class  reflect.Value
	before reflect.Value
	after  reflect.Value
}
type option struct{}
type routerOption struct {
	prefix         string
	methodList     []string
	middlewareList []*middleware
}
type helperRouter struct {
	ANY     string
	GET     string
	POST    string
	OPTIONS string
	Option  *option
}
type router struct {
	prefix         string
	methodList     []string
	path           string
	controller     reflect.Value
	action         reflect.Value
	parameter      []reflect.Value
	middlewareList []*middleware
	regexp         map[string]string
}

var (
	Router           = newHelperRouter()
	routerMethodList = []string{Router.ANY, Router.GET, Router.POST}
	routerList       []*router
	routerOptionList []*routerOption
)

func newMiddleware(newMiddleware interface{}) *middleware {
	class := reflect.New(reflect.ValueOf(newMiddleware).Type().Elem())
	return &middleware{
		class:  class,
		before: class.MethodByName("Before"),
		after:  class.MethodByName("After"),
	}
}
func newOption() *option {
	return &option{}
}
func newRouterOption() *routerOption {
	return &routerOption{
		methodList:     []string{},
		middlewareList: []*middleware{},
	}
}
func newHelperRouter() *helperRouter {
	return &helperRouter{
		GET:     "GET",
		POST:    "POST",
		OPTIONS: "OPTIONS",
		Option:  newOption(),
	}
}
func newRouter() *router {
	return &router{
		methodList:     []string{},
		parameter:      []reflect.Value{},
		middlewareList: []*middleware{},
		regexp:         map[string]string{},
	}
}
func (this *helperRouter) Any(path string, controller interface{}, action string) *router {
	return this.MethodList([]string{Router.ANY}, path, controller, action)
}
func (this *helperRouter) Get(path string, controller interface{}, action string) *router {
	return this.MethodList([]string{Router.GET}, path, controller, action)
}
func (this *helperRouter) Post(path string, controller interface{}, action string) *router {
	return this.MethodList([]string{Router.POST}, path, controller, action)
}
func (this *helperRouter) Method(method string, path string, controller interface{}, action string) *router {
	if !_slice.In(method, routerMethodList) {
		interceptor.Insure(false).
			Message("方法不支持").
			Data(map[string]interface{}{"method": method}).
			Do()
	}
	return this.MethodList([]string{method}, path, controller, action)
}
func (this *helperRouter) MethodList(methodList []string, path string, controller interface{}, action string) *router {
	r := newRouter()
	r.methodList = append(r.methodList, methodList...)
	for _, o := range routerOptionList {
		r.prefix += o.prefix
		for _, middleware := range o.middlewareList {
			r.middlewareList = append(r.middlewareList, middleware)
		}
		r.methodList = append(r.methodList, o.methodList...)
	}
	r.path = path
	r.controller = reflect.New(reflect.ValueOf(controller).Type().Elem())
	r.action = r.controller.MethodByName(action)
	if r.action.Type().NumIn() > 0 {
		r.parameter = append(r.parameter, reflect.ValueOf(this))
	}
	routerList = append(routerList, r)
	return r
}
func (this *helperRouter) Group(routerOption *routerOption, group func()) {
	group()
	if len(routerOptionList) > 0 {
		routerOptionList = routerOptionList[0 : len(routerOptionList)-1]
	}
}
func (this *routerOption) Prefix(prefix string) *routerOption {
	this.prefix += prefix
	return this
}
func (this *routerOption) Method(method string) *routerOption {
	this.methodList = append(this.methodList, method)
	return this
}
func (this *routerOption) MethodList(methodList []string) *routerOption {
	this.methodList = append(this.methodList, methodList...)
	return this
}
func (this *routerOption) Middleware(middleware interface{}) *routerOption {
	this.middlewareList = append(this.middlewareList, newMiddleware(middleware))
	return this
}
func (this *routerOption) MiddlewareList(middlewareList []interface{}) *routerOption {
	for _, middleware := range middlewareList {
		this.middlewareList = append(this.middlewareList, newMiddleware(middleware))
	}
	return this
}
func (this *option) Prefix(prefix string) *routerOption {
	routerOption := newRouterOption()
	routerOption.prefix = prefix
	routerOptionList = append(routerOptionList, routerOption)
	return routerOption
}
func (this *option) Method(method string) *routerOption {
	routerOption := newRouterOption()
	routerOption.methodList = []string{method}
	routerOptionList = append(routerOptionList, routerOption)
	return routerOption
}
func (this *option) MethodList(methodList []string) *routerOption {
	routerOption := newRouterOption()
	routerOption.methodList = methodList
	routerOptionList = append(routerOptionList, routerOption)
	return routerOption
}
func (this *option) Middleware(middleware interface{}) *routerOption {
	routerOption := newRouterOption()
	routerOption.middlewareList = append(routerOption.middlewareList, newMiddleware(middleware))
	routerOptionList = append(routerOptionList, routerOption)
	return routerOption
}
func (this *option) MiddlewareList(middlewareList []interface{}) *routerOption {
	routerOption := newRouterOption()
	for _, middleware := range middlewareList {
		routerOption.middlewareList = append(routerOption.middlewareList, newMiddleware(middleware))
	}
	routerOptionList = append(routerOptionList, routerOption)
	return routerOption
}
func (this *router) Prefix(prefix string) *router {
	this.prefix += prefix
	return this
}
func (this *router) Middleware(middleware interface{}) *router {
	this.middlewareList = append(this.middlewareList, newMiddleware(middleware))
	return this
}
func (this *router) MiddlewareList(middlewareList []interface{}) *router {
	for _, middleware := range middlewareList {
		this.middlewareList = append(this.middlewareList, newMiddleware(middleware))
	}
	return this
}
func (this *router) Regexp(regexp map[string]string) *router {
	this.regexp = regexp
	return this
}
