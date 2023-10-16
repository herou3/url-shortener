package server

import (
	"github.com/herou3/url-shortener/internal/app/internal/services"
	"net/http"
	"regexp"
)

// Server is a server with all the batteries included
type Server struct {
	Mux *http.ServeMux
}

// Init returns new server instance
func Init() *Server {
	server := &Server{
		Mux: http.NewServeMux(),
	}

	server.Mux.HandleFunc(`/`, pathHandler())

	return server
}

func pathHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/" {
			services.HandleCreateShortURL(writer, request)
			return
		}
		status, _ := regexp.MatchString("([/][a-zA-Z0-9]{8}$)", request.URL.Path)
		if status && len(request.URL.Path) == 9 {
			services.HandleGetFullURL(writer, request)
			return
		}

		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}
