package main

import (
	"log"
	"net/http"
	"os"

	service "github.com/hyphengolang/flyio/internal/service/http"
)

// fly.io requires port 8080
var port = os.Getenv("PORT")

func init() {
	if port == "" {
		port = "8080"
	}
}

func main() {
	srv := service.NewService()

	log.Println("Listening on port", port)
	log.Fatalln(http.ListenAndServe(":"+port, srv))
}
