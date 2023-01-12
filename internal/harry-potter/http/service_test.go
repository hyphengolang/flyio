package http_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/hyphengolang/prelude/testing/is"

	hpSrv "github.com/hyphengolang/flyio/internal/harry-potter/http"
)

var srv *httptest.Server

func init() {
	mux := chi.NewRouter()

	mux.Mount("/characters", hpSrv.NewService())

	srv = httptest.NewServer(mux)
}

func TestService(t *testing.T) {
	is := is.New(t)

	t.Cleanup(func() { srv.Close() })

	//NOTE - Date obj in JS The string format should be: YYYY-MM-DDTHH:mm:ss.sssZ
	// Also accepts YYYY-MM or YYYY

	t.Run("create a new character", func(t *testing.T) {
		payload := `
		{
			"name": "Harry Potter",
			"blood": "muggle-born",
			"species": "human",
			"born": "1980-07-31",
			"quote": "I don't go looking for trouble. Trouble usually finds me.",
			"imgUrl": "https://i.pinimg.com/originals/58/39/61/5839613bc887946211e72778d01da05f.jpg"
		}`

		res, err := http.Post(srv.URL+"/characters", "application/json", strings.NewReader(payload))
		is.NoErr(err)                                // no request error
		is.Equal(res.StatusCode, http.StatusCreated) // post character
	})

	t.Run("get all characters", func(t *testing.T) {
		res, err := http.Get(srv.URL + "/characters")
		is.NoErr(err) // no request error

		is.Equal(res.StatusCode, http.StatusOK) // get all characters
	})
}
