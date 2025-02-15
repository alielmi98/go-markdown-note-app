package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/alielmi98/go-markdown-note-app/services"
)

// NoteHandler handles HTTP requests related to notes.
type NoteHandler struct {
	noteService *services.NoteService
}

// NewNoteHandler creates a new NoteHandler with the given NoteService.
func NewNoteHandler(service *services.NoteService) *NoteHandler {
	return &NoteHandler{noteService: service}
}

// UploadHandler handles the uploading of markdown files.
func (h *NoteHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the uploaded file from the request.
	_, file, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}

	// Check if the uploaded file has a .md extension.
	if strings.ToLower(strings.Split(file.Filename, ".")[1]) != "md" {
		http.Error(w, "File must be a markdown file", http.StatusBadRequest)
		return
	}

	// Call the service to upload the note.
	_, err = h.noteService.UploadNote(file, h.noteService.UploadDir)
	if err != nil {
		http.Error(w, "Failed to upload file", http.StatusInternalServerError)
		return
	}

	// Log the successful upload and respond with a success message.
	log.Printf("%s uploaded successfully", file.Filename)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "File uploaded successfully"})
}
