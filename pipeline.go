package pioneer

import (
	"net/http"

	"golang.org/x/net/context"
)

type Pipeline struct {
	handler Handler
	plug   []Plugger
}

func (p *Pipeline) HTTPHandler() http.Handler {
	var h Handler = p.handler
	for i := len(p.plug) - 1; i >= 0; i-- {
		h = p.plug[i].Plug(h)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		h.ServeHTTP(ctx, w, r)
	})
}

func (p *Pipeline) Handler() Handler {
	var h Handler = p.handler
	for i := len(p.plug) - 1; i >= 0; i-- {
		h = p.plug[i].Plug(h)
	}
	return h
}

func (p *Pipeline) Plug(plug ...Plugger) *Pipeline {
	p.plug = append(p.plug, plug...)
	return p
}

func (p *Pipeline) Plugs() []Plugger {
	return p.plug
}

func (p *Pipeline) SetHandler(h Handler) *Pipeline {
	p.handler = h
	return p
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		handler: HandlerFunc(func(context.Context, http.ResponseWriter, *http.Request) {}),
	}
}
