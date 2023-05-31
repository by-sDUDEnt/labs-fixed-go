package rooms

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go-labs-game-platform/internal/httpserver/httphelpers"
)

func (h *Handlers) GetByID(w http.ResponseWriter, r *http.Request) {
	roomIDStr := mux.Vars(r)["room_id"]

	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		httphelpers.WriteError(w, r, 0, err)
		return
	}

	list, err := h.srv.GetByID(r.Context(), roomID)
	if err != nil {
		httphelpers.WriteError(w, r, 0, err)
		return
	}

	httphelpers.WriteJSON(w, r, http.StatusOK, list)
}
