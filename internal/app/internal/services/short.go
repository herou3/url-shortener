package services

import (
	"io"
	"net/http"
	"regexp"
)

var uh = UrlHandler{}

// HandleCreateShortUrl Call method for generate new short url
func HandleCreateShortUrl(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	fullUrl, err := io.ReadAll(request.Body)
	isMatch, err := regexp.MatchString("((http|https)://)(www.)?[a-zA-Z0-9@:%._+~#?&/=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%._+~#?&/=]*)", string(fullUrl))
	if !isMatch {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	short, err := uh.CreateShortUrl(string(fullUrl))

	_, err = response.Write([]byte(short))
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	response.WriteHeader(http.StatusCreated)
}

// HandleGetFullUrl Call method for get actual full address for short email
func HandleGetFullUrl(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		response.WriteHeader(http.StatusBadRequest)
		return
	}
	id := request.URL.Path[1:len(request.URL.Path)]

	fu, err := uh.GetFullUrl(id)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	response.Header().Add("Location", fu)
	response.WriteHeader(http.StatusTemporaryRedirect)
}
