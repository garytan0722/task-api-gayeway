package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
}

func (rw *ResponseRecorder) WriteHeader(code int) {
	rw.StatusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func RequestLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recorder := &ResponseRecorder{
			ResponseWriter: w,
			StatusCode:     http.StatusOK, // default to 200
		}

		start := time.Now()
		next.ServeHTTP(recorder, r)
		duration := time.Since(start)

		entry := log.WithFields(log.Fields{
			"method":   r.Method,
			"path":     r.URL.Path,
			"status":   recorder.StatusCode,
			"duration": duration.Seconds(),
		})

		if recorder.StatusCode >= 200 && recorder.StatusCode < 300 {
			entry.Debug("OK")
		} else {
			entry.Error("Request failed")
		}
	}
}
