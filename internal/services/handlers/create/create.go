package handlers

import (
	"io"
	"net/http"
	"regexp"

	"github.com/herou3/url-shortener/internal/services/context"
)

func CreateShortURLHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	fullURL, _ := io.ReadAll(request.Body)
	// This regexp describes a restrictions for correct url, if url from request doesn't have correct form this call will return an error
	isMatch, _ := regexp.MatchString("((http|https)://)(www.)?[a-zA-Z0-9@:%._+~#?&/=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%._+~#?&/=]*)", string(fullURL))
	if !isMatch {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	short, _ := context.GetUH().CreateShortURL(string(fullURL))

	response.WriteHeader(http.StatusCreated)
	_, err := response.Write([]byte(short))
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
}
