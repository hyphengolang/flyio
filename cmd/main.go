package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	// hpSrv "github.com/hyphengolang/flyio/internal/harry-potter/http"
	baseSrv "github.com/hyphengolang/flyio/internal/service/http"
)

// fly.io requires port 8080
var port = os.Getenv("PORT")

func init() {
	if port == "" {
		port = "8080"
	}
}

func main() {
	// default
	mux := chi.NewRouter()
	mux.Mount("/", baseSrv.NewService())
	// mux.Mount("/characters", hpSrv.NewService())

	log.Println("Listening on port", port)
	log.Fatalln(http.ListenAndServe(":"+port, mux))
}
