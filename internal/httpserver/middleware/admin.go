package middleware

import (
	"net/http"
	"strings"

	"go-labs-game-platform/internal/config"
	"go-labs-game-platform/internal/httpserver/httphelpers"
	"go-labs-game-platform/internal/models"
)

func Admin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authStr := r.Header.Get("Authorization")
		parts := strings.Split(authStr, " ")
		if len(parts) != 2 || parts[0] != "Bearer" || parts[1] != config.Get().Security.AdminSecret {
			httphelpers.WriteError(w, r, 0, models.ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
