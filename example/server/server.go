package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/iotalabs/pioneer"
	"github.com/iotalabs/pioneer/plugs/body"
	"github.com/iotalabs/pioneer/plugs/close"
	"github.com/iotalabs/pioneer/plugs/limit"
	"github.com/iotalabs/pioneer/plugs/router"
	"github.com/iotalabs/pioneer/plugs/static"
	"github.com/iotalabs/pioneer/utils"
)

func helloRoute(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	params, _ := router.FetchParams(ctx)
	utils.DumpJSON(
		w,
		map[string]interface{}{
			"status": "success",
			"msg":    "Hello, " + params.ByName("name") + "!",
		},
	)
}

func slowRoute(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	utils.DumpJSON(w, `Im slow!`)
}

func newRouter() *router.Router {
	r := router.New()
	r.Get("/api/hello/:name", pioneer.HandlerFunc(helloRoute))
	r.Post("/api/hello/:name", pioneer.HandlerFunc(helloRoute))
	r.Get("/api/slow", pioneer.HandlerFunc(slowRoute))
	return r
}

func main() {
	p := pioneer.NewPipeline()

	p.Plug(limit.New(1, time.Second))
	p.Plug(close.New(func(r *http.Request) {
		fmt.Printf("CLOSED: %#v\n", r.RemoteAddr)
	}))
	p.Plug(
		static.New(
			static.Dir("./static"),
			static.Prefix("public"),
		),
	)
	p.Plug(body.New(func(r *http.Request, err error) {
		fmt.Println("Error: ", err)
	}))
	p.Plug(newRouter())

	http.ListenAndServe(":8080", p.HTTPHandler())
}
