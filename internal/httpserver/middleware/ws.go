package middleware

import "net/http"

func SetTokenFromQuery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if token := r.URL.Query().Get("token"); token != "" {
			r.Header.Set("Authorization", "Bearer "+token)
		}

		next.ServeHTTP(w, r)
	})
}
