package tools

import (
	"fmt"
	"testing"
)

func TestStringWithCharset(t *testing.T) {
	tests := []struct {
		name    string
		charset string
		size    int
		want    string
	}{
		{
			name:    "Check generation for simple of String with one char",
			charset: "a",
			size:    4,
			want:    "aaaa",
		},
		{
			name:    "Check generation string whose size equals 0 chars",
			charset: "x",
			size:    0,
			want:    "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gs := StringWithCharset(test.size, test.charset)
			if len(gs) == test.size &&
				gs == test.want {
				return
			}
			t.Errorf("StringWithCharset() = %s, want %s", fmt.Sprint(gs), test.want)
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{
			name: "Check string with definite size",
			size: 7,
			want: 7,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gs := String(test.size)
			if len(gs) == test.size {
				return
			}
			t.Errorf("StringWithCharset() = %d, want %d", len(gs), test.want)
		})
	}
}
