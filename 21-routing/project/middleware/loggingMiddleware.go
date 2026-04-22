package middleware

import (
	"fmt"
	"net/http"
	"time"
	"project/utils"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if rw.status == 0 {
		rw.status = 200
	}
	return rw.ResponseWriter.Write(b)
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		rw := &responseWriter{ResponseWriter: w}

		defer func() {
			if err := recover(); err != nil {
				res, _ := utils.CreateResponse(false, "Internal Server Error", nil)
				utils.WriteJSON(rw, 500, res)
			}
		}()

		start := time.Now()
		timestamp := start.Format(time.RFC3339)

		if r.URL.Path != "/favicon.ico" {
			fmt.Printf("Incoming: %s %s\n", r.Method, r.URL.Path)
		}

		next.ServeHTTP(rw, r)

		if r.URL.Path != "/favicon.ico" {
			fmt.Printf("[%s] %s %s → %d → %s\n",
				timestamp,
				r.Method,
				r.URL.Path,
				rw.status,
				time.Since(start),
			)
		}
	})
}
