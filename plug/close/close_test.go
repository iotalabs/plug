package close

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/iotalabs/pioneer"
	"golang.org/x/net/context"
)

func TestClose(t *testing.T) {
	close := New(nil)

	dur := 2 * time.Second
	pipe := pioneer.NewPipeline()
	pipe.Plug(close)
	pipe.Plug(pioneer.HandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		select {
		case <-ctx.Done():
		case <-time.After(dur + time.Second):
			t.Error("Should cancel\n")
		}
	}))

	ts := httptest.NewServer(pipe.HTTPHandler())
	defer ts.Close()

	go http.Get(ts.URL)
	time.Sleep(dur - time.Second)
	ts.CloseClientConnections()
}
