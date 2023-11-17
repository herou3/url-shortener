package handlers

import "testing"

func TestHandleCreateShortURL(t *testing.T) {
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

		})
	}
}
