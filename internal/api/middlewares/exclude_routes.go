package middlewares

import (
	"net/http"
	"strings"
)

func MiddlewaresExcludePaths(middleware func(http.Handler) http.Handler, excludedPaths ...string) func(http.Handler) http.Handler {
	// fmt.Println("MiddlwareExcludePaths initialized")
	return func(next http.Handler) http.Handler {
		// fmt.Println("MiddlwareExcludePaths Running")
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, path := range excludedPaths {
				if strings.HasPrefix(r.URL.Path, path) {
					next.ServeHTTP(w, r)
					return
				}
			}
			middleware(next).ServeHTTP(w, r)
			// fmt.Println("Sent response from MiddlwareExcludePaths")
		})
	}
}
