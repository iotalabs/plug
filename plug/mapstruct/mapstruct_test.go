package mapstruct

import (
	"bytes"
	"net/http"
	"testing"

	"golang.org/x/net/context"

	"github.com/iotalabs/pioneer"
	"github.com/stretchr/testify/assert"
)

func TestPlug(t *testing.T) {
	p := New(nil)
	pipe := pioneer.NewPipeline().Plug(p)
	pipe.SetHandler(
		pioneer.HandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
			var d struct {
				Name string
				Age  int
			}
			err := Decode(ctx, &d)
			assert.Nil(t, err)
			assert.Equal(t, d.Name, "Alex")
			assert.Equal(t, d.Age, 27)
		}),
	)
	dogPayload := `{"name": "Alex", "age": 27}`
	r, _ := http.NewRequest("GET", "/", bytes.NewBufferString(dogPayload))
	pipe.HTTPHandler().ServeHTTP(nil, r)
}
