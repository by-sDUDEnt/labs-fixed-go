package user

import (
	"net/http"

	"go-labs-game-platform/internal/ctxpkg"
	"go-labs-game-platform/internal/httpserver/httphelpers"
)

func (h *Handlers) GetMe(w http.ResponseWriter, r *http.Request) {
	userID := ctxpkg.GetUserID(r.Context())

	user, err := h.srv.ByID(r.Context(), userID)
	if err != nil {
		httphelpers.WriteError(w, r, 0, err)
		return
	}

	httphelpers.WriteJSON(w, r, http.StatusOK, user)
}
