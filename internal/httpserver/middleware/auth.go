package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"go-labs-game-platform/internal/ctxpkg"
	"go-labs-game-platform/internal/httpserver/httphelpers"
	"go-labs-game-platform/internal/models"
	"go-labs-game-platform/internal/services/auth"
)

func Authorization(srv auth.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authStr := r.Header.Get("Authorization")
			parts := strings.Split(authStr, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				httphelpers.WriteError(w, r, 0, models.ErrUnauthorized)
				return
			}

			plaintext := parts[1]

			r, err := processAuth(r, srv, plaintext)
			if err != nil {
				httphelpers.WriteError(w, r, 0, err)
				return
			}

			if err = srv.UpdateLastVisitedAt(r.Context(), plaintext); err != nil {
				httphelpers.WriteError(w, r, 0, fmt.Errorf("update last visited: %w", err))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func processAuth(r *http.Request, srv auth.Service, tokenText string) (*http.Request, error) {
	token, err := srv.GetToken(r.Context(), tokenText)
	if err != nil {
		return r, err
	}

	if !token.IsValid() {
		return r, models.ErrUnauthorized
	}

	switch token.Scope {
	case models.ScopeSessionUser:
		ctx := ctxpkg.SetUserID(r.Context(), token.UserID)
		return r.WithContext(ctx), nil
	}

	return r, models.ErrUnauthorized
}
