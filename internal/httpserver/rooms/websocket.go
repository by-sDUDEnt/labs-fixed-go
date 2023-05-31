package rooms

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"go-labs-game-platform/internal/ctxpkg"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handlers) WebSocket(w http.ResponseWriter, r *http.Request) {
	userID := ctxpkg.GetUserID(r.Context())
	roomIDStr := mux.Vars(r)["room_id"]

	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		log.Println(err)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// check if room exists
	room, err := h.srv.GetByID(r.Context(), roomID)
	if err != nil {
		log.Println(err)
		return
	}

	if room == nil {
		log.Println("room == nil")
		return
	}

	go func() {
		for {
			// read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			if msgType != websocket.TextMessage {
				log.Println("msgType != websocket.TextMessage")
				return
			}

			// print message to console
			log.Printf("%s sent: %s", conn.RemoteAddr(), string(msg))

			if err = h.srv.HandlePacket(r.Context(), userID, roomID, msg); err != nil {
				log.Println(err)
				return
			}
		}
	}()

	packets, err := h.srv.ListenPackets(r.Context(), roomID, userID)
	if err != nil {
		log.Println(err)
		return
	}

	for packet := range packets {
		if err = conn.WriteJSON(packet.Payload); err != nil {
			log.Println(err)
			return
		}
	}
}
