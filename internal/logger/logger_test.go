package logger

import (
	"go.uber.org/zap"
	"testing"
)

func TestInitializeLogger(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Check to return logger with 'debug' level logging",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var log = zap.NewNop()
			err := InitializeLogger("info")
			if err != nil {
				t.Errorf("InitializeLogger() was returned err - %t, wanted - %s", err, "info")
			}
			if log == nil {
				t.Errorf("InitializeLogger() was returned nil, wanted - %s", "info")
			}
		})
	}
}
