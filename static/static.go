package static

import (
	"net/http"
	"path"
	"strings"

	"github.com/iotalabs/pioneer"
	"golang.org/x/net/context"
)

type static struct {
	dir    string
	prefix string

	fileHandler http.Handler

	next pioneer.Handler
}

// Dir set real directory to serve static files from
func Dir(path string) func(*static) {
	return func(s *static) {
		s.dir = path
	}
}

// Prefix set prefix to url to serve static files
func Prefix(prefix string) func(*static) {
	return func(s *static) {
		s.prefix = path.Join("/", prefix)
	}
}

// New create a new static file server
// Default:
//  dir: ./static
//  prefix: /public
func New(setups ...func(*static)) pioneer.Plugger {
	s := &static{
		dir:    "./static",
		prefix: "/public",
	}

	for _, setup := range setups {
		setup(s)
	}

	s.fileHandler = http.StripPrefix(
		s.prefix,
		http.FileServer(
			http.Dir(s.dir),
		),
	)

	return s
}

// ServeHTTP implements pioneer.Handler
func (s *static) ServeHTTP(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	if strings.HasPrefix(req.URL.Path, s.prefix) {
		s.fileHandler.ServeHTTP(res, req)
		return
	}
	s.next.ServeHTTP(ctx, res, req)
}

// Plug implements pioneer.Plugger
func (s *static) Plug(h pioneer.Handler) pioneer.Handler {
	s.next = h
	return s
}
