package handlers

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleCreateShortURL(t *testing.T) {
	type header struct {
		headerName  string
		headerValue string
	}
	type want struct {
		code   int
		header header
	}
	tests := []struct {
		name         string
		beforeAction bool
		body         string
		want         want
	}{
		{
			name:         "Check to create shortUrl for https",
			beforeAction: false,
			body:         "https://ya.ru/",
			want: want{
				code: 201,
			},
		},
		{
			name:         "Check to create shortUrl for http",
			beforeAction: false,
			body:         "http://test.ru/",
			want: want{
				code: 201,
			},
		},
		{
			name:         "Check to create error in process creation client",
			beforeAction: true,
			body:         "golang.ru",
			want: want{
				code: 400,
			},
		},
		{
			name:         "Check to create error in process creation client, request with empty body",
			beforeAction: true,
			body:         "",
			want: want{
				code: 400,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(test.body)))
			w := httptest.NewRecorder()
			HandleCreateShortURL(w, request)
			res := w.Result()
			assert.Equal(t, test.want.code, res.StatusCode)
			err := res.Body.Close()
			if err != nil {
				return
			}
		})
	}
}
