package server

import (
	"github.com/go-chi/chi/v5"
	create "github.com/herou3/url-shortener/internal/services/handlers/create"
	"github.com/herou3/url-shortener/internal/services/handlers/get"
)

// Server is a server with all the batteries included
type Server struct {
	Mux *chi.Mux
}

// Init returns new server instance
func Init() *Server {
	server := &Server{
		Mux: chi.NewRouter(),
	}

	server.Mux.Post("/", create.HandleCreateShortURL)
	server.Mux.Get("/{id}", get.HandleGetFullURL)

	return server
}
