package rooms

import (
	"net/http"

	"go-labs-game-platform/internal/ctxpkg"
	"go-labs-game-platform/internal/httpserver/httphelpers"
)

func (h *Handlers) Create(w http.ResponseWriter, r *http.Request) {
	userID := ctxpkg.GetUserID(r.Context())

	room, err := h.srv.Create(r.Context(), userID)
	if err != nil {
		httphelpers.WriteError(w, r, 0, err)
		return
	}

	httphelpers.WriteJSON(w, r, http.StatusCreated, room)
}
