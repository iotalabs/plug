package body

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/iotalabs/pioneer"

	"golang.org/x/net/context"
)

func init() {
	log.Println("This plug is deprecated. Please use mapstruct plug.")
}

type body struct {
	onError func(r *http.Request, err error)
	next    pioneer.Handler
}

type readCloser struct {
	rc  io.ReadCloser
	bp  *body
	buf buffer
}

func newBody(onError func(r *http.Request, err error)) *body {
	return &body{onError: onError}
}

// New create new body plug
func New(onError func(r *http.Request, err error)) pioneer.Plugger {
	return newBody(onError)
}

// ErrBodyNotPlugged body plug not plugged
var ErrBodyNotPlugged = errors.New("body not plugged")

// DecodeJSON decode json to interface from body
func DecodeJSON(ctx context.Context, v interface{}) error {
	data, ok := FetchBody(ctx)

	if !ok {
		return ErrBodyNotPlugged
	}

	return json.Unmarshal(data, v)
}

var ctxBodyKey uint8

// FetchBody fetch request body from context
func FetchBody(ctx context.Context) ([]byte, bool) {
	body, ok := ctx.Value(&ctxBodyKey).(buffer)
	return []byte(body), ok
}

func (bp *body) Plug(h pioneer.Handler) pioneer.Handler {
	bp.next = h
	return bp
}

func (bp *body) ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if _, ok := FetchBody(ctx); !ok {
		buf := getBuffer()
		defer buf.free()

		if _, err := io.Copy(&buf, r.Body); err != nil {
			if bp.onError != nil {
				bp.onError(r, err)
			}
		} else {
			ctx = context.WithValue(ctx, &ctxBodyKey, buf)
		}
	}

	bp.next.ServeHTTP(ctx, w, r)
	return
}
