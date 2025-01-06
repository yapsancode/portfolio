# Portfolio Website

A modern portfolio website built with:
- Go
- HTMX
- Tailwind CSS
- MySQL

## Features
- Interactive quiz system
- Contact form
- Ask Me Anything section
- Audio welcome message
- Dynamic content loading with HTMX

## Setup
1. Clone the repository
2. Set up MySQL database
3. Configure database connection in `internal/database/mysql.go`
4. Run `go mod tidy` to install dependencies
5. Run `go run cmd/server/main.go` to start the server

## Project Structure
```tree
portfolio/
├── cmd/
│   └── server/
├── internal/
│   ├── config/
│   ├── handlers/
│   ├── models/
│   └── database/
├── static/
└── templates/
