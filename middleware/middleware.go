package middleware

import "net/http"

func Content(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "Aplication/json")
		next.ServeHTTP(w, r)
	})
}
