package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
	"go-labs-game-platform/internal/services/room"
)

func main() {
	token := os.Args[1]
	//	connect to ws server
	roomID := "9f08466c-a3f8-4815-b5a2-5e601513090b"
	url := "ws://localhost:8080/api/v1/game/" + roomID + "/ws?token=" + token

	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	fmt.Println("connected")

	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			log.Printf("recv: %s", message)
		}
	}()

	// read from stdin and send to ws server

	for {
		var t string
		_, err := fmt.Scanln(&t)
		if err != nil {
			log.Println("scan:", err)
			return
		}

		switch t {
		case "s":
			marshal, err := json.Marshal(room.ClientStartPacket{
				Packet: room.Packet{
					Type: room.PacketTypeClientStart,
				},
			})
			if err != nil {
				fmt.Println(err)
				return
			}

			t = string(marshal)
		case "m":
			marshal, err := json.Marshal(room.ClientMovePacket{
				Packet: room.Packet{
					Type: room.PacketTypeClientMove,
				},

				// {"type": 0, "position": 1}
				Position: 1,
			})
			if err != nil {
				fmt.Println(err)
				return
			}

			t = string(marshal)
		default:
			fmt.Println("unknown command")
			continue
		}

		if err = c.WriteMessage(websocket.TextMessage, []byte(t)); err != nil {
			log.Println("write:", err)
			return
		}
	}
}
