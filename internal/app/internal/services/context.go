package services

import (
	"errors"
	"fmt"
	"github.com/herou3/url-shortener/internal/app/internal/services/tools"
)

// Errors
var (
	errFullURLIsEmpty = errors.New("property in body for fullEmail should no be empty")
	errShortURLEmpty  = errors.New("query property for shortEmail should no be empty")
	errURLIsNotFound  = errors.New("data hasn't information about this url")
)

type URLCreatorGetter interface {
	CreateShortURL(url string) (string, error)
	GetFullURL(shortUrl string) (string, error)
}

type URLHandler struct {
	storage UrlsStorage
}

type URLDetail struct {
	fullURL  string
	shortURL string
}

func (h *URLHandler) CreateShortURL(fullURL string) (string, error) {
	if len(fullURL) == 0 {
		return "", fmt.Errorf("fullURL is empty %s", errFullURLIsEmpty)
	}
	shortUrl, err := h.storage.CreateShortURL(fullURL)
	if err != nil {
		return "", fmt.Errorf("create url: %w", err)
	}

	return shortUrl, nil
}

func (h *URLHandler) GetFullURL(shortURL string) (string, error) {
	if len(shortURL) == 0 {
		return "", fmt.Errorf("fullURL is empty %s", errShortURLEmpty)
	}
	fullURL, err := h.storage.GetFullURL(shortURL)
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
	urls map[string]URLDetail
}

func (us *UrlsStorage) CreateShortURL(fullURL string) (string, error) {
	isNotFound := true
	shortURLId := ""
	for isNotFound {
		shortURLId = tools.String(8)
		_, err := us.GetFullURL(shortURLId)

		if err != nil {
			isNotFound = false
		}
	}

	urlDetail := URLDetail{
		fullURL:  fullURL,
		shortURL: shortURLId,
	}

	if us.urls == nil {
		us.urls = make(map[string]URLDetail)
	}
	us.urls[urlDetail.shortURL] = urlDetail

	return shortURLId, nil
}

func (us *UrlsStorage) GetFullURL(shortURL string) (string, error) {
	urlDetail, ok := us.urls[shortURL]
	if !ok {
		return "", errURLIsNotFound
	}

	return urlDetail.fullURL, nil
}
