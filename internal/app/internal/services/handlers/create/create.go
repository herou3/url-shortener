package handlers

import (
	"github.com/herou3/url-shortener/internal/app/internal/services/context"
	"io"
	"net/http"
	"regexp"
)

var uh = context.GetUH()

func HandleCreateShortURL(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	fullURL, _ := io.ReadAll(request.Body)
	isMatch, _ := regexp.MatchString("((http|https)://)(www.)?[a-zA-Z0-9@:%._+~#?&/=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%._+~#?&/=]*)", string(fullURL))
	if !isMatch {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	short, _ := uh.CreateShortURL(string(fullURL))

	_, err := response.Write([]byte(short))
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	response.WriteHeader(http.StatusCreated)
}
