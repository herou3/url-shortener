package server

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Check to init Server",
			want: "fsd",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := Init()
			assert.Contains(t, reflect.TypeOf(server).String(), test.want)
		})
	}
}
