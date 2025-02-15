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

// Delete Note deletes the specified note file from the server.
func (s *NoteService) DeleteNote(filename string) error {
	// Lock the mutex to ensure thread safety
	s.mu.Lock()
	defer s.mu.Unlock()

	// Create the file path
	filePath := filepath.Join(s.UploadDir, filename)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return err
	}

	// Delete the file
	if err := os.Remove(filePath); err != nil {
		return err
	}

	return nil

}

func (s *NoteService) GetNoteList() ([]string, error) {
	// Lock the mutex to ensure thread safety
	s.mu.Lock()
	defer s.mu.Unlock()

	// Open the directory
	dir, err := os.ReadDir(s.UploadDir)
	if err != nil {
		return nil, err
	}

	// Create a slice to store the file names
	var files []string

	// Iterate over the directory entries
	for _, entry := range dir {
		// Check if the entry is a file
		if entry.Type().IsRegular() {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}

func (s *NoteService) GetNoteContent(filename string) ([]byte, error) {
	// Lock the mutex to ensure thread safety
	s.mu.Lock()
	defer s.mu.Unlock()

	// Create the file path
	filePath := filepath.Join(s.UploadDir, filename)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the content of the file
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}
