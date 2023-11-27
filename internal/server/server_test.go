package server

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Check to init Server",
			want: "Server",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := Init()
			assert.Contains(t, reflect.TypeOf(server).String(), test.want)
		})
	}
}
