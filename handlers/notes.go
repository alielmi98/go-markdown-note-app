package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/alielmi98/go-markdown-note-app/services"
	"github.com/alielmi98/go-markdown-note-app/utils"
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

func (h *NoteHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the filename from the request URL.
	filename := strings.TrimPrefix(r.URL.Path, "/api/notes/")
	if filename == "" {
		http.Error(w, "Invalid filename", http.StatusBadRequest)
		return
	}

	// Call the service to delete the note.
	err := h.noteService.DeleteNote(filename)
	if err != nil {
		http.Error(w, "Failed to delete file", http.StatusInternalServerError)
		log.Printf("Failed to delete file: %s", err)
		return
	}

	// Log the successful deletion and respond with a success message.
	log.Printf("%s deleted successfully", filename)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "File deleted successfully"})
}

func (h *NoteHandler) ListNotesHandler(w http.ResponseWriter, r *http.Request) {
	// Call the service to list all notes.
	notes, err := h.noteService.GetNoteList()
	if err != nil {
		http.Error(w, "Failed to list notes", http.StatusInternalServerError)
		log.Printf("Failed to list notes: %s", err)
		return
	}

	// Respond with the list of notes.
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}

func (h *NoteHandler) RenderToHtmlHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the filename from the request URL.
	filename := strings.TrimPrefix(r.URL.Path, "/api/notes/")
	if filename == "" {
		http.Error(w, "Invalid filename", http.StatusBadRequest)
		return
	}

	content, err := h.noteService.GetNoteContent(filename)
	if err != nil {
		http.Error(w, "Failed to get file content", http.StatusInternalServerError)
		log.Printf("Failed to get file content: %s", err)
		return
	}

	// Parse the template file
	tmpl, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		http.Error(w, "Failed to parse template file", http.StatusInternalServerError)
		log.Printf("Failed to parse template file: %s", err)
		return
	}

	htmlContent := utils.MdToHtml(content)

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, map[string]interface{}{
		"PageTitle": filename,
		"Content":   template.HTML(string(htmlContent)),
	})

}
