package context

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/herou3/url-shortener/internal/config"
	"github.com/herou3/url-shortener/internal/services/tools"
)

// Errors
var (
	errFullURLIsEmpty = errors.New("property in body for fullEmail should no be empty")
	errShortURLEmpty  = errors.New("query property for shortEmail should no be empty")
	errURLIsNotFound  = errors.New("data hasn't information about this url")
)

type URLCreatorGetter interface {
	CreateShortURL(url string) (string, error)
	GetFullURL(shortURL string) (string, error)
}

type URLHandler struct {
	Storage UrlsStorage
}

var myHandler *URLHandler

func GetUH() *URLHandler {
	if myHandler == nil {
		myHandler = &URLHandler{}
	}

	return myHandler
}

type URLDetail struct {
	FullURL  string
	ShortURL string
}

func (h *URLHandler) CreateShortURL(fullURL string) (string, error) {
	if len(fullURL) == 0 {
		return "", fmt.Errorf("fullURL is empty %s", errFullURLIsEmpty)
	}
	shortURL, err := h.Storage.CreateShortURL(fullURL)
	if err != nil {
		return "", fmt.Errorf("create url: %w", err)
	}

	return shortURL, nil
}

func (h *URLHandler) GetFullURL(shortURL string) (string, error) {
	if len(shortURL) == 0 {
		return "", fmt.Errorf("fullURL is empty %s", errShortURLEmpty)
	}
	fullURL, err := h.Storage.GetFullURL(shortURL)
	if err != nil {
		return "", fmt.Errorf("create url: %w", err)
	}
	if fullURL == "" {
		return "", errURLIsNotFound
	}

	return fullURL, nil
}

// Methods for db
type UrlsStorage struct {
	Urls map[string]URLDetail
}

func (us *UrlsStorage) CreateShortURL(fullURL string) (string, error) {
	isNotFound := true
	idForShortURL := ""
	if fullURL == "" {
		return "", fmt.Errorf("fullURL is empty %s", errFullURLIsEmpty)
	}
	for isNotFound {
		idForShortURL = tools.String(8)
		_, err := us.GetFullURL(idForShortURL)

		if err != nil {
			isNotFound = false
		}
	}

	urlDetail := URLDetail{
		FullURL:  fullURL,
		ShortURL: idForShortURL,
	}

	if us.Urls == nil {
		us.Urls = make(map[string]URLDetail)
	}
	us.Urls[urlDetail.ShortURL] = urlDetail

	isHasHTTP, _ := regexp.MatchString("((http|https)://)", config.GetConf().ShortURL)

	if isHasHTTP {
		return config.GetConf().ShortURL + "/" + idForShortURL, nil
	} else {
		return "https" + config.GetConf().ShortURL + "/" + idForShortURL, nil
	}
}

func (us *UrlsStorage) GetFullURL(shortURL string) (string, error) {
	urlDetail, ok := us.Urls[shortURL]
	if !ok {
		return "", errURLIsNotFound
	}

	return urlDetail.FullURL, nil
}
