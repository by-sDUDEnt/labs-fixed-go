package rooms

import (
	"net/http"

	"go-labs-game-platform/internal/httpserver/httphelpers"
)

func (h *Handlers) GetList(w http.ResponseWriter, r *http.Request) {
	list, err := h.srv.GetList(r.Context())
	if err != nil {
		httphelpers.WriteError(w, r, 0, err)
		return
	}

	httphelpers.WriteJSON(w, r, http.StatusOK, list)
}
