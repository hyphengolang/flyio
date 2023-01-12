package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	h "github.com/hyphengolang/prelude/http"
)

type Router interface {
	chi.Router

	Respond(w http.ResponseWriter, r *http.Request, data any, status int)
	Decode(w http.ResponseWriter, r *http.Request, data any) error
}

type router struct {
	chi.Router
}

func (m *router) Respond(w http.ResponseWriter, r *http.Request, data any, status int) {
	h.Respond(w, r, data, status)
}

func (m *router) Decode(w http.ResponseWriter, r *http.Request, data any) error {
	return h.Decode(w, r, data)
}

func NewRouter() Router {
	m := &router{
		Router: chi.NewRouter(),
	}
	// do nothings
	return m
}
