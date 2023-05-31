package auth

import (
	"net/http"

	"go-labs-game-platform/internal/httpserver/httphelpers"
	"go-labs-game-platform/internal/models"
)

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var (
		req = models.RegisterCredentials{}
	)

	if err := httphelpers.ReadBody(r, &req); err != nil {
		httphelpers.WriteError(w, r, 0, err)
		return
	}

	result, err := h.authSrv.Register(r.Context(), &req)
	if err != nil {
		httphelpers.WriteError(w, r, 0, err)
		return
	}

	httphelpers.WriteJSON(w, r, http.StatusOK, result)
}
