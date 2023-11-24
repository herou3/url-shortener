package get

import (
	"github.com/herou3/url-shortener/internal/services/context"
	"net/http"
)

var uh = context.GetUH()

// Get full url handler
func HandleGetFullURL(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	id := request.URL.Path[1:len(request.URL.Path)]

	fu, err := uh.GetFullURL(id)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	response.Header().Add("Location", fu)
	response.WriteHeader(http.StatusTemporaryRedirect)
}
