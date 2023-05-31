package rooms

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go-labs-game-platform/internal/ctxpkg"
	"go-labs-game-platform/internal/httpserver/httphelpers"
)

func (h *Handlers) Leave(w http.ResponseWriter, r *http.Request) {
	userID := ctxpkg.GetUserID(r.Context())
	roomIDStr := mux.Vars(r)["room_id"]

	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		httphelpers.WriteError(w, r, 0, err)
		return
	}

	if err = h.srv.Leave(r.Context(), userID, roomID); err != nil {
		httphelpers.WriteError(w, r, 0, err)
		return
	}

	httphelpers.WriteJSON(w, r, http.StatusOK, nil)
}
