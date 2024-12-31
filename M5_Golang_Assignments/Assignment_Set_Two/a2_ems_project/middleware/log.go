package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Create a response writer wrapper to capture status code
		ww := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(ww, r)

		duration := time.Since(startTime)
		fmt.Printf("Request Method: %s | Path: %s | Status: %d | Duration: %v\n",
			r.Method, r.URL.Path, ww.statusCode, duration)
	})
}

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (ww *responseWriterWrapper) WriteHeader(code int) {
	ww.statusCode = code
	ww.ResponseWriter.WriteHeader(code)
}
