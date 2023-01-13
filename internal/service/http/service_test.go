package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/hyphengolang/prelude/testing/is"

	baseSrv "github.com/hyphengolang/flyio/internal/service/http"
)

var srv *httptest.Server

func init() {
	mux := chi.NewRouter()

	mux.Mount("/", baseSrv.NewService())

	srv = httptest.NewServer(mux)
}

func TestService(t *testing.T) {
	is := is.New(t)

	t.Cleanup(func() { srv.Close() })

	t.Run("make a get request with Name query", func(t *testing.T) {
		res, err := srv.Client().Get(srv.URL + "/?name=John")
		is.NoErr(err)                           // no request error
		is.Equal(res.StatusCode, http.StatusOK) // success
		defer res.Body.Close()

		type response struct {
			Message string `json:"message"`
		}

		var v response
		err = json.NewDecoder(res.Body).Decode(&v)
		is.NoErr(err)                      // no decode error
		is.Equal(v.Message, "Hello John!") // check response message
	})
}
