package context

import (
	"github.com/herou3/url-shortener/internal/app/internal/config"
	"github.com/stretchr/testify/assert"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestGetUH(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "Check to return UH",
			want: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			event := GetUH()
			if event == nil {
				currentStatus := event == nil
				t.Errorf("GetUH() was returned nil - %t, wanted - %t", currentStatus, test.want)
			}
			assert.Contains(t, reflect.TypeOf(GetUH()).String(), "URLHandler")
		})
	}
}

func TestUrlsStorage_GetFullURL(t *testing.T) {
	type want struct {
		fullURL  string
		shortURL string
	}
	tests := []struct {
		name         string
		beforeAction bool
		want         want
	}{
		{
			name:         "Check to return FullURL by shortURL",
			beforeAction: true,
			want: want{
				fullURL:  "https://ya.ru/",
				shortURL: "aaaaaaa",
			},
		},
		{
			name: "Check to return error with description when urls hadn't been in context",
			want: want{
				fullURL:  "",
				shortURL: "bbbbbbb",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.beforeAction {
				GetUH().Storage.Urls = make(map[string]URLDetail)
				GetUH().Storage.Urls[test.want.shortURL] = URLDetail{
					FullURL:  test.want.fullURL,
					ShortURL: test.want.shortURL,
				}
			}
			result, _ := GetUH().Storage.GetFullURL(test.want.shortURL)
			assert.Equal(t, result, test.want.fullURL)
		})
	}
}

func TestUrlsStorage_CreateShortURL(t *testing.T) {
	type want struct {
		fullURL      string
		sizeShortURL int
	}
	tests := []struct {
		name         string
		beforeAction bool
		want         want
	}{
		{
			name:         "Check to return fullURL by shortURL",
			beforeAction: true,
			want: want{
				fullURL:      "https://ya.ru/",
				sizeShortURL: 8,
			},
		},
		{
			name: "Check to return error with description when urls hadn't been in context",
			want: want{
				fullURL:      "",
				sizeShortURL: 0,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := GetUH().Storage.CreateShortURL(test.want.fullURL)
			strs := strings.Split(result, "/")
			if len(strs) > 2 {
				assert.Equal(t, len(strs[2]), test.want.sizeShortURL)
			}
		})
	}
}

func TestURLHandler_GetFullURL(t *testing.T) {
	type want struct {
		fullURL  string
		shortURL string
	}
	tests := []struct {
		name         string
		beforeAction bool
		want         want
	}{
		{
			name:         "Check to return fullURL by shortURL",
			beforeAction: true,
			want: want{
				fullURL:  "https://ya.ru/",
				shortURL: "aaaaaaa",
			},
		},
		{
			name: "Check to return error with description when urls hadn't been in context",
			want: want{
				fullURL:  "",
				shortURL: "bbbbbbb",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.beforeAction {
				GetUH().Storage.Urls = make(map[string]URLDetail)
				GetUH().Storage.Urls[test.want.shortURL] = URLDetail{
					FullURL:  test.want.fullURL,
					ShortURL: test.want.shortURL,
				}
			}
			result, _ := GetUH().GetFullURL(test.want.shortURL)
			assert.Equal(t, result, test.want.fullURL)
		})
	}
}

func TestURLHandler_CreateShortURL(t *testing.T) {
	type want struct {
		fullURL      string
		sizeShortURL int
	}
	tests := []struct {
		name         string
		beforeAction bool
		want         want
	}{
		{
			name:         "Check to return fullURL by shortURL",
			beforeAction: true,
			want: want{
				fullURL:      "https://ya.ru/",
				sizeShortURL: 25,
			},
		},
		{
			name: "Check to return error with description when urls hadn't been in context",
			want: want{
				fullURL:      "",
				sizeShortURL: 0,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := GetUH().CreateShortURL(test.want.fullURL)
			isHasHttp, _ := regexp.MatchString("((http|https)://)", config.GetConf().ShortURL)
			if isHasHttp {
				assert.Equal(t, len(result), test.want.sizeShortURL)
			}
		})
	}
}
