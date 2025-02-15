package main

import (
	"log"
	"net/http"

	"github.com/alielmi98/go-markdown-note-app/routers"
)

func main() {
	r := routers.NewRouter()

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
