package http

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	h "github.com/hyphengolang/prelude/http"
)

type Router interface {
	chi.Router

	SetLocation(w http.ResponseWriter, r *http.Request, id string)
	Respond(w http.ResponseWriter, r *http.Request, data any, status int)
	Decode(w http.ResponseWriter, r *http.Request, data any) error

	Log(v ...any)
	Logf(format string, v ...any)
}

type router struct {
	chi.Router
}

func (m *router) SetLocation(w http.ResponseWriter, r *http.Request, id string) {
	path := r.URL.Path
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	w.Header().Add("Location", "//"+r.Host+path+id)
}

func (m *router) Respond(w http.ResponseWriter, r *http.Request, data any, status int) {
	h.Respond(w, r, data, status)
}

func (m *router) Decode(w http.ResponseWriter, r *http.Request, data any) error {
	return h.Decode(w, r, data)
}

func (m *router) Log(v ...any) {
	log.Println(v...)
}

func (m *router) Logf(format string, v ...any) {
	log.Printf(format, v...)
}

func NewRouter() Router {
	m := &router{
		Router: chi.NewRouter(),
	}
	// do nothings
	return m
}
