package timeout

import (
	"net/http"
	"time"

	"github.com/iotalabs/pioneer"
	"golang.org/x/net/context"
)

// timeout plug
// if duration <= 0, no timeout
type timeout struct {
	next     pioneer.Handler
	Duration time.Duration
}

// New create a new timeout plug
func New(d time.Duration) pioneer.Plugger {
	return newTimeout(d)
}

func newTimeout(d time.Duration) *timeout {
	return &timeout{Duration: d}
}

// Plug implements pioneer.Plugger interface
func (to *timeout) Plug(h pioneer.Handler) pioneer.Handler {
	to.next = h
	return to
}

// ServeHTTP implements pioneer.Handler interface
func (to *timeout) ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if to.Duration > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, to.Duration)
		defer cancel()
	}
	to.next.ServeHTTP(ctx, w, r)
}
