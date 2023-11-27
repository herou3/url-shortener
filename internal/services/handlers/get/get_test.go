package get

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/herou3/url-shortener/internal/services/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleGetFullURL(t *testing.T) {
	type header struct {
		headerName  string
		headerValue string
	}
	type want struct {
		code        int
		header      header
		otherMethod bool
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
			name:         "Check to the book was located in context",
			beforeAction: true,
			want: want{
				code: 307,
				header: header{
					headerName:  "Location",
					headerValue: "https://ya.ru/",
				},
			},
		},
		{
			name:         "Check to the book was located in context",
			beforeAction: true,
			want: want{
				code: 400,
				header: header{
					headerName:  "Location",
					headerValue: "https://ya.ru/",
				},
				otherMethod: true,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defaultClientID := "msymbol"
			if test.beforeAction {
				context.GetUH().Storage.Urls = make(map[string]context.URLDetail)
				context.GetUH().Storage.Urls[defaultClientID] = context.URLDetail{
					FullURL:  test.want.header.headerValue,
					ShortURL: defaultClientID,
				}
			}
			request := httptest.NewRequest(http.MethodGet, "/"+defaultClientID, nil)
			if test.want.otherMethod {
				request = httptest.NewRequest(http.MethodPost, "/"+defaultClientID, nil)
			}
			w := httptest.NewRecorder()
			HandleGetFullURL(w, request)
			res := w.Result()
			assert.Equal(t, test.want.code, res.StatusCode)
			if test.beforeAction && !test.want.otherMethod {
				_, err := io.ReadAll(res.Body)
				require.NoError(t, err)
				assert.Equal(t, test.want.header.headerValue, res.Header.Get(test.want.header.headerName))

				context.GetUH().Storage.Urls = make(map[string]context.URLDetail)
			}
			err := res.Body.Close()
			if err != nil {
				return
			}
		})
	}
}
