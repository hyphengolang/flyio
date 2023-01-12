package http

import (
	"fmt"
	"net/http"

	"github.com/hyphengolang/flyio/internal/common"
	router "github.com/hyphengolang/flyio/internal/http"
)

type service struct {
	mux router.Router
}

func (s *service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// Handle's default routes
func NewService() http.Handler {
	s := &service{
		mux: router.NewRouter(),
	}
	s.routes()
	return s
}

func (s *service) routes() {
	s.mux.Get("/*", s.handleRoot())

}

func (s *service) handleRoot() http.HandlerFunc {
	type response struct {
		Message string `json:"message"`
		Meta    struct {
			Hostname  string `json:"hostname"`
			IPAddress string `json:"ipAddress"`
		} `json:"meta"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		// get Name from request url query
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "World"
		}

		v := response{
			Message: fmt.Sprintf("Hello %s!", name),
			Meta: struct {
				Hostname  string `json:"hostname"`
				IPAddress string `json:"ipAddress"`
			}{
				Hostname:  common.GetHostname(),
				IPAddress: common.GetIPAddress(),
			},
		}

		s.mux.Respond(w, r, v, http.StatusOK)
	}
}
