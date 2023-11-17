package get

import (
	"github.com/herou3/url-shortener/internal/app/internal/services/context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetFullURL(t *testing.T) {
	type want struct {
		code   int
		site   string
		header string
	}
	tests := []struct {
		name         string
		beforeAction bool
		want         want
	}{
		{
			name:         "Check to the book was not located in context",
			beforeAction: false,
			want: want{
				code: 400,
			},
		},
		{
			name:         "Check to the book wasn located in context",
			beforeAction: true,
			want: want{
				site: "https://yandex.ru/",
				code: 307,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defaultClientId := "msymbol"
			if test.beforeAction {
				context.GetUH().Storage.Urls = make(map[string]context.URLDetail)
				context.GetUH().Storage.Urls[defaultClientId] = context.URLDetail{
					FullURL:  test.want.site,
					ShortURL: defaultClientId,
				}
			}
			request := httptest.NewRequest(http.MethodGet, "/"+defaultClientId, nil)
			w := httptest.NewRecorder()
			HandleGetFullURL(w, request)
			res := w.Result()
			assert.Equal(t, test.want.code, res.StatusCode)
			if test.beforeAction {
				context.GetUH().Storage.Urls = make(map[string]context.URLDetail)
			}
		})
	}
}
