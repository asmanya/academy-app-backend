package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	// fmt.Println("ResponseTime middleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("responseTime middlware being returned =========XXXXXXXXXX==========")
		start := time.Now()

		// create a custom response writer to capture the code
		wrappedWriter := &responseWriter{ResponseWriter: w, status: http.StatusOK}

		duration := time.Since(start)
		wrappedWriter.Header().Set("X-Response-Time", duration.String())
		next.ServeHTTP(wrappedWriter, r)

		// calculate the duration
		duration = time.Since(start)

		// log the request details
		fmt.Printf("Method: %s. URL: %s, Status: %d, Duration: %v\n", r.Method, r.URL, wrappedWriter.status, duration.String())

		// fmt.Println("Sent response from responseTime middleware")
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
