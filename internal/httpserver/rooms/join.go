package rooms

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go-labs-game-platform/internal/ctxpkg"
	"go-labs-game-platform/internal/httpserver/httphelpers"
)

func (h *Handlers) Join(w http.ResponseWriter, r *http.Request) {
	userID := ctxpkg.GetUserID(r.Context())
	// get path param mux
	roomIDStr := mux.Vars(r)["room_id"]

	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		httphelpers.WriteError(w, r, 0, err)
		return
	}

	room, err := h.srv.Join(r.Context(), userID, roomID)
	if err != nil {
		httphelpers.WriteError(w, r, 0, err)
		return
	}

	httphelpers.WriteJSON(w, r, http.StatusCreated, room)
}
