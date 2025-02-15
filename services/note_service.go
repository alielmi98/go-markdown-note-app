package services

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"sync"
)

// NoteService provides methods to handle note operations.
type NoteService struct {
	UploadDir string
	mu        sync.Mutex
}

// NewNoteService creates a new NoteService with the specified upload directory.
func NewNoteService(uploadDir string) *NoteService {
	return &NoteService{UploadDir: uploadDir}
}

// UploadNote processes the uploaded markdown file and saves it to the server.
func (s *NoteService) UploadNote(file *multipart.FileHeader, uploadDir string) (string, error) {
	// Lock the mutex to ensure thread safety
	s.mu.Lock()
	defer s.mu.Unlock()

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Create the destination file
	dstPath := filepath.Join(uploadDir, file.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Read the content from the uploaded file and write it to the destination file
	data, err := io.ReadAll(src)
	if err != nil {
		return "", err
	}

	_, err = dst.Write(data)
	if err != nil {
		return "", err
	}

	return dstPath, nil
}
