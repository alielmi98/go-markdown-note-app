package routers

import (
	"github.com/alielmi98/go-markdown-note-app/internal/handlers"
	"github.com/alielmi98/go-markdown-note-app/internal/services"

	"github.com/gorilla/mux"
)

// NewRouter creates a new mux.Router and sets up the routes for the API.
func NewRouter() *mux.Router {
	// Initialize the NoteService with the directory for uploads.
	service := services.NewNoteService("./uploads")

	// Create a new NoteHandler with the NoteService.
	handler := handlers.NewNoteHandler(service)

	// Create a new mux.Router.
	r := mux.NewRouter()

	// Define routes for the API.
	r.HandleFunc("/api/notes", handler.UploadHandler).Methods("POST")
	r.HandleFunc("/api/notes/{filename}", handler.DeleteHandler).Methods("DELETE")
	r.HandleFunc("/api/notes", handler.ListNotesHandler).Methods("GET")
	r.HandleFunc("/api/notes/{filename}", handler.RenderToHtmlHandler).Methods("GET")
	r.HandleFunc("/api/notes/check/{filename}", handler.CheckGrammarHandler).Methods("GET")

	return r
}
