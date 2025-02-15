# go-markdown-note-app

A RESTful API built with Go that allows users to upload and manage markdown notes. Users can upload markdown files, check the grammar, save the note, and retrieve the HTML rendered version of the note. The app includes endpoints for grammar checking, note saving, note listing, and rendered note retrieval. This project demonstrates how to handle file uploads, parse and render markdown files, and perform grammar checks in a Go-based RESTful API.

This project is part of the [roadmap.sh](https://roadmap.sh) backend intermediate level projects. You can find the project details [here](https://roadmap.sh/projects/markdown-note-taking-app).

## Project Structure
```bash
go-markdown-note-app/
├── internal/
│   ├── handlers/
│   │   ├── notes.go
│   ├── routers/
│   │   ├── router.go
│   ├── services/
│   │   ├── note_service.go
│   ├── utils/
│   │   ├── grammar_checker.go
│   │   ├── markdown_renderer.go
├── templates/
│   ├── index.tmpl
├── uploads/
│   └── (uploaded markdown files)
├── .gitignore
├── .gitattributes
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── main.go
```
## Getting Started

1. Clone the repository:
   ```sh
   git clone https://github.com/alielmi98/go-markdown-note-app.git
   cd go-markdown-note-app

2. Install dependencies:
   ```sh
    go mod tidy

3. Install dependencies:
   ```sh
    go run main.go


## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License - see the LICENSE file for details.