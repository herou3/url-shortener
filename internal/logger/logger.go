package logger

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Header string

type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}

type (
	responseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

var Log = zap.NewNop()

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode // захватываем код статуса
}

// InitializeLogger logger with identify level for logging
func InitializeLogger(level string) error {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}
	cfg := zap.NewProductionConfig()
	cfg.Level = lvl
	zl, err := cfg.Build()
	if err != nil {
		return err
	}

	Log = zl
	return nil
}

func WithLogging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		start := time.Now()
		uri := req.RequestURI
		method := req.Method

		responseData := &responseData{
			status: 0,
			size:   0,
		}
		lw := loggingResponseWriter{
			ResponseWriter: res, // встраиваем оригинальный http.ResponseWriter
			responseData:   responseData,
		}

		// point for making original request
		h.ServeHTTP(&lw, req)

		duration := time.Since(start)

		Log.Info("got incoming HTTP request",
			zap.String("uri", uri),
			zap.String("method", method),
			zap.Float64("duration", duration.Seconds()),
		)
		Log.Info("got information from service",
			zap.Int("status", responseData.status),
			zap.Int("size", responseData.size),
		)
	})
}
