package context

import (
	"github.com/stretchr/testify/assert"
	"reflect"
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
		fullUrl  string
		shortUrl string
	}
	tests := []struct {
		name         string
		beforeAction bool
		want         want
	}{
		{
			name:         "Check to return fullUrl by shortUrl",
			beforeAction: true,
			want: want{
				fullUrl:  "https://ya.ru/",
				shortUrl: "aaaaaaa",
			},
		},
		{
			name: "Check to return error with description when urls hadn't been in context",
			want: want{
				fullUrl:  "",
				shortUrl: "bbbbbbb",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.beforeAction {
				GetUH().Storage.Urls = make(map[string]URLDetail)
				GetUH().Storage.Urls[test.want.shortUrl] = URLDetail{
					FullURL:  test.want.fullUrl,
					ShortURL: test.want.shortUrl,
				}
			}
			result, _ := GetUH().Storage.GetFullURL(test.want.shortUrl)
			assert.Equal(t, result, test.want.fullUrl)
		})
	}
}

func TestUrlsStorage_CreateShortURL(t *testing.T) {
	type want struct {
		fullUrl      string
		sizeShortUrl int
	}
	tests := []struct {
		name         string
		beforeAction bool
		want         want
	}{
		{
			name:         "Check to return fullUrl by shortUrl",
			beforeAction: true,
			want: want{
				fullUrl:      "https://ya.ru/",
				sizeShortUrl: 8,
			},
		},
		{
			name: "Check to return error with description when urls hadn't been in context",
			want: want{
				fullUrl:      "",
				sizeShortUrl: 0,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := GetUH().Storage.CreateShortURL(test.want.fullUrl)
			assert.Equal(t, len(result), test.want.sizeShortUrl)
		})
	}
}

func TestURLHandler_GetFullURL(t *testing.T) {
	type want struct {
		fullUrl  string
		shortUrl string
	}
	tests := []struct {
		name         string
		beforeAction bool
		want         want
	}{
		{
			name:         "Check to return fullUrl by shortUrl",
			beforeAction: true,
			want: want{
				fullUrl:  "https://ya.ru/",
				shortUrl: "aaaaaaa",
			},
		},
		{
			name: "Check to return error with description when urls hadn't been in context",
			want: want{
				fullUrl:  "",
				shortUrl: "bbbbbbb",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.beforeAction {
				GetUH().Storage.Urls = make(map[string]URLDetail)
				GetUH().Storage.Urls[test.want.shortUrl] = URLDetail{
					FullURL:  test.want.fullUrl,
					ShortURL: test.want.shortUrl,
				}
			}
			result, _ := GetUH().GetFullURL(test.want.shortUrl)
			assert.Equal(t, result, test.want.fullUrl)
		})
	}
}

func TestURLHandler_CreateShortURL(t *testing.T) {
	type want struct {
		fullUrl      string
		sizeShortUrl int
	}
	tests := []struct {
		name         string
		beforeAction bool
		want         want
	}{
		{
			name:         "Check to return fullUrl by shortUrl",
			beforeAction: true,
			want: want{
				fullUrl:      "https://ya.ru/",
				sizeShortUrl: 8,
			},
		},
		{
			name: "Check to return error with description when urls hadn't been in context",
			want: want{
				fullUrl:      "",
				sizeShortUrl: 11,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := GetUH().CreateShortURL(test.want.fullUrl)
			assert.Equal(t, len(result), test.want.sizeShortUrl)
		})
	}
}
