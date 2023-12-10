package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/herou3/url-shortener/internal/services/handlers/create"
	"github.com/herou3/url-shortener/internal/services/handlers/get"
)

// Server is a server with all the batteries included
type Server struct {
	Mux *chi.Mux
}

// CreateServerInstance returns new server instance
func CreateServerInstance() *Server {
	server := &Server{
		Mux: chi.NewRouter(),
	}

	server.Mux.Post("/", handlers.CreateShortURLHandler)
	server.Mux.Get("/{id}", get.CallFullURLHandler)

	return server
}
