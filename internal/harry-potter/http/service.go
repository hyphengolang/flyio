package http

import (
	"net/http"
	"time"

	router "github.com/hyphengolang/flyio/internal/http"
)

type Option func(*service)

// TODO

type service struct {
	mux router.Router
}

func (s *service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// Handle's all /characters routes
func NewService(opts ...Option) http.Handler {
	s := &service{
		mux: router.NewRouter(),
	}

	for _, opt := range opts {
		opt(s)
	}

	s.routes()
	return s
}

func (s *service) routes() {
	s.mux.Get("/", s.handleList())
	s.mux.Get("/{id}", s.handleGet())
	s.mux.Post("/", s.handleCreate())
	s.mux.Put("/{id}", s.handleUpdate())
	s.mux.Delete("/{id}", s.handleDelete())
}

func (s *service) handleList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.mux.Respond(w, r, "list characters", http.StatusNotImplemented)
	}
}

func (s *service) handleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.mux.Respond(w, r, "get characters by ID", http.StatusNotImplemented)
	}
}

func (s *service) handleCreate() http.HandlerFunc {
	type request struct {
		Name    string     `json:"name"`
		Blood   string     `json:"blood"`
		Species string     `json:"species"`
		Born    *time.Time `json:"born"`
		Quote   string     `json:"quote"`
		ImgURL  string     `json:"imgUrl"`
	}

	decode := func(w http.ResponseWriter, r *http.Request, character any) error {
		var req request
		if err := s.mux.Decode(w, r, &req); err != nil {
			return err
		}

		return nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var c any
		if err := decode(w, r, c); err != nil {
			s.mux.Respond(w, r, err, http.StatusBadRequest)
			return
		}

		s.mux.Respond(w, r, "create character", http.StatusOK)
	}
}

func (s *service) handleUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.mux.Respond(w, r, "update character", http.StatusNotImplemented)
	}
}

func (s *service) handleDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.mux.Respond(w, r, "delete character", http.StatusNotImplemented)
	}
}
