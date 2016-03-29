package router

import "github.com/iotalabs/pioneer"

type SubRouter struct {
	prefix   string
	pipeline *pioneer.Pipeline
	father   *Router
}

func (sr *SubRouter) SubRouter(prefix string) *SubRouter {
	return sr.father.subRouter(sr, prefix)
}

func (sr *SubRouter) handle(method string, path string, h pioneer.Handler) {
	sr.father.handle(sr, method, path, h)
}

func (sr *SubRouter) Plug(plug ...pioneer.Plugger) {
	sr.father.plug(sr, plug...)
}

func (sr *SubRouter) Get(path string, h pioneer.Handler) {
	sr.handle(GET, path, h)
}

func (sr *SubRouter) Post(path string, h pioneer.Handler) {
	sr.handle(POST, path, h)
}

func (sr *SubRouter) Put(path string, h pioneer.Handler) {
	sr.handle(PUT, path, h)
}

func (sr *SubRouter) Patch(path string, h pioneer.Handler) {
	sr.handle(PATCH, path, h)
}

func (sr *SubRouter) Options(path string, h pioneer.Handler) {
	sr.handle(OPTIONS, path, h)
}

func (sr *SubRouter) Delete(path string, h pioneer.Handler) {
	sr.handle(DELETE, path, h)
}
