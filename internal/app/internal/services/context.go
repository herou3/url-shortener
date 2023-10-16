package services

import (
	"errors"
	"fmt"
	"github.com/herou3/url-shortener/internal/app/internal/services/tools"
)

// Errors
var (
	errFullUrlIsEmpty = errors.New("property in body for fullEmail should no be empty")
	errShortUrlEmpty  = errors.New("query property for shortEmail should no be empty")
	errUrlIsNotFound  = errors.New("data hasn't information about this url")
)

type UrlCreatorGetter interface {
	CreateShortUrl(url string) (string, error)
	GetFullUrl(shortUrl string) (string, error)
}

type UrlHandler struct {
	storage UrlsStorage
}

type UrlDetail struct {
	fullUrl  string
	shortUrl string
}

func (h *UrlHandler) CreateShortUrl(fullUrl string) (string, error) {
	if len(fullUrl) == 0 {
		return "", fmt.Errorf("fullUrl is empty", errFullUrlIsEmpty)
	}
	shortUrl, err := h.storage.CreateShortUrl(fullUrl)
	if err != nil {
		return "", fmt.Errorf("create url: %w", err)
	}

	return shortUrl, nil
}

func (h *UrlHandler) GetFullUrl(shortUrl string) (string, error) {
	if len(shortUrl) == 0 {
		return "", fmt.Errorf("fullUrl is empty", errShortUrlEmpty)
	}
	fullUrl, err := h.storage.GetFullUrl(shortUrl)
	if err != nil {
		return "", fmt.Errorf("create url: %w", err)
	}
	if fullUrl == "" {
		return "", errUrlIsNotFound
	}

	return fullUrl, nil
}

// Methods for db
type UrlsStorage struct {
	urls map[string]UrlDetail
}

func (us *UrlsStorage) CreateShortUrl(fullUrl string) (string, error) {
	isNotFound := true
	shortUrlId := ""
	for isNotFound {
		shortUrlId = tools.String(8)
		_, err := us.GetFullUrl(shortUrlId)

		if err != nil {
			isNotFound = false
		}
	}

	urlDetail := UrlDetail{
		fullUrl:  fullUrl,
		shortUrl: shortUrlId,
	}

	if us.urls == nil {
		us.urls = make(map[string]UrlDetail)
	}
	us.urls[urlDetail.shortUrl] = urlDetail

	return shortUrlId, nil
}

func (us *UrlsStorage) GetFullUrl(shortUrl string) (string, error) {
	urlDetail, ok := us.urls[shortUrl]
	if !ok {
		return "", errUrlIsNotFound
	}

	return urlDetail.fullUrl, nil
}
