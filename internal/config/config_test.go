package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestURLHandler_CreateShortURL(t *testing.T) {
	type want struct {
		host string
		url  string
	}
	tests := []struct {
		name           string
		isEmptyRequest bool
		want           want
	}{
		{
			name:           "Check to return default config",
			isEmptyRequest: true,
			want: want{
				host: "localhost:8080",
				url:  "local:8989",
			},
		},
		{
			name: "Check to return config with customs values",
			want: want{
				host: "local:7070",
				url:  "http://test:1010/",
			},
		},
		{
			name: "Check to return config without information about port",
			want: want{
				host: "newHost",
				url:  "http://test/",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			conf := GetConf()
			if test.isEmptyRequest {
				conf.SetConf(nil)
			} else {
				conf.SetConf(map[string]string{"host": test.want.host, "shortURL": test.want.url})
			}
			assert.Equal(t, conf.HOST, test.want.host)
			assert.Equal(t, conf.ShortURL, test.want.url)
		})
	}
}
