package middleware

import (
	"fmt"
	"net/http"
	"time"

	"dadandev.com/wa-engine/pkg/utils"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(w http.ResponseWriter, r *http.Request) {
			if err := recover(); err != nil {
				utils.ErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
			}
		}(w, r)
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: 200}
		h.ServeHTTP(rec, r)
		duration := time.Since(start)
		fmt.Printf("[%s] %s | %d | %v\n",
			r.Method,
			r.URL.Path,
			rec.status,
			duration,
		)
	})
}
