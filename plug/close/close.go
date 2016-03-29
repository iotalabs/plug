package close

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/iotalabs/pioneer"
)

// Closer client connection close plug
type closer struct {
	next    pioneer.Handler
	onClose func(r *http.Request)
}

// New create closer plug, optional pass in a onClose func
func New(onClose func(r *http.Request)) pioneer.Plugger {
	return &closer{onClose: onClose}
}

// Plug implements pioneer.Plugger interface
func (c *closer) Plug(h pioneer.Handler) pioneer.Handler {
	c.next = h
	return c
}

// ServeHTTP implements pioneer.Handler interface
func (c *closer) ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// Cancel the context if the client closes the connection
	if cn, ok := w.(http.CloseNotifier); ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		go func() {
			<-cn.CloseNotify()
			cancel()
			if c.onClose != nil {
				c.onClose(r)
			}
		}()
	}

	c.next.ServeHTTP(ctx, w, r)
}
