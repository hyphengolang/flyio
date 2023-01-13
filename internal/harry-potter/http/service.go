package http

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	intern "github.com/hyphengolang/flyio/internal"
	hp "github.com/hyphengolang/flyio/internal/harry-potter"
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

var characters = map[uuid.UUID]hp.Character{}

func (s *service) handleList() http.HandlerFunc {
	type response struct {
		Characters []hp.Character `json:"characters"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var chars []hp.Character
		for _, c := range characters {
			chars = append(chars, c)
		}

		s.mux.Respond(w, r, response{Characters: chars}, http.StatusOK)
	}
}

func (s *service) handleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.mux.Respond(w, r, "get characters by ID", http.StatusNotImplemented)
	}
}

func (s *service) handleCreate() http.HandlerFunc {
	type request struct {
		Name    string        `json:"name"`
		Blood   hp.BloodTyp   `json:"blood"`
		Species hp.SpeciesTyp `json:"species"`
		Born    *intern.Time  `json:"born"`
		Quote   string        `json:"quote"`
		ImgURL  string        `json:"imgUrl"`
	}

	decode := func(w http.ResponseWriter, r *http.Request, character *hp.Character) error {
		var req request
		if err := s.mux.Decode(w, r, &req); err != nil {
			return fmt.Errorf("harry-potter: invalid request: %w", err)
		}

		imgUrl, err := url.Parse(req.ImgURL)
		if err != nil {
			return fmt.Errorf("harry-potter: invalid image url: %w", err)
		}

		*character = hp.Character{
			ID:      uuid.New(),
			Name:    req.Name,
			Blood:   req.Blood,
			Species: req.Species,
			Born:    req.Born,
			Quote:   req.Quote,
			ImgURL:  imgUrl,
		}

		return nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var c hp.Character
		if err := decode(w, r, &c); err != nil {
			s.mux.Respond(w, r, err, http.StatusBadRequest)
			return
		}

		characters[c.ID] = c

		s.mux.SetLocation(w, r, c.ID.String())
		s.mux.Respond(w, r, "create character", http.StatusCreated)
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
