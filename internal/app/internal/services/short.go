package services

import (
	"io"
	"net/http"
	"regexp"
)

var uh = URLHandler{}

// Generate new short url
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
