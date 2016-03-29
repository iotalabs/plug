package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/iotalabs/pioneer"
	"golang.org/x/net/context"
)

func main() {
	p := pioneer.NewPipeline()

	p.Plug(pioneer.HandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Hello")
	}))

	start := time.Now()
	p.HTTPHandler().ServeHTTP(nil, nil)
	fmt.Println(time.Now().Sub(start))
}
