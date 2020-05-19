package middleware

import "net/http"

func ValidContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Invalid content-type", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
