package timeout

import (
	"net/http"
	"testing"
	"time"

	"github.com/iotalabs/pioneer"
	"golang.org/x/net/context"
)

func TestTimeout(t *testing.T) {
	to := newTimeout(1 * time.Second)

	pipe := pioneer.NewPipeline()
	pipe.Plug(to)
	pipe.Plug(pioneer.HandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		select {
		case <-ctx.Done():
		case <-time.After(to.Duration + time.Second):
			t.Error("Shoud timeout\n")
		}
	}))
	pipe.HTTPHandler().ServeHTTP(nil, nil)

	to = newTimeout(0)
	pipe = pioneer.NewPipeline()
	pipe.Plug(to)
	pipe.Plug(pioneer.HandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		select {
		case <-ctx.Done():
			t.Error("Shoud not timeout\n")
		case <-time.After(time.Second):
		}
	}))
	pipe.HTTPHandler().ServeHTTP(nil, nil)
}
