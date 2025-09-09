package middlewares

import (
	"compress/gzip"
	"net/http"
	"strings"
)

func Compression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if the client acceptss gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		// set the response header
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()

		// write the responseWriter
		w = &gzipResponseWriter{ResponseWriter: w, writer: gz}

		next.ServeHTTP(w, r)
		// fmt.Println("Sent response from compression middleware")
	})
}

// gzipResponseWriter wraps http.ResponseWriter to write gzipped repsonses
type gzipResponseWriter struct {
	http.ResponseWriter
	writer *gzip.Writer
}

func (g *gzipResponseWriter) Write(b []byte) (int, error) {
	return g.writer.Write(b)
}

// this middleware helps in compressing the response sent to the client
