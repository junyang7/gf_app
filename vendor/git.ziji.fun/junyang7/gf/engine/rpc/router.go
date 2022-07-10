package rpc

import (
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
	middlewareList []*middleware
}
type helperRouter struct {
	Option *option
}
type router struct {
	prefix         string
	path           string
	controller     reflect.Value
	action         reflect.Value
	parameter      []reflect.Value
	middlewareList []*middleware
	regexp         map[string]string
}

var (
	Router           = newHelperRouter()
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
		middlewareList: []*middleware{},
	}
}
func newHelperRouter() *helperRouter {
	return &helperRouter{
		Option: newOption(),
	}
}
func newRouter() *router {
	return &router{
		parameter:      []reflect.Value{},
		middlewareList: []*middleware{},
		regexp:         map[string]string{},
	}
}
func (this *helperRouter) Any(path string, controller interface{}, action string) *router {
	r := newRouter()
	for _, o := range routerOptionList {
		r.prefix += o.prefix
		for _, middleware := range o.middlewareList {
			r.middlewareList = append(r.middlewareList, middleware)
		}
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
