package room

import "fmt"

var (
	ErrRoomIsFull                = fmt.Errorf("room is full")
	ErrRoomAlreadyContainsPlayer = fmt.Errorf("room already contains player")
	ErrRoomDoesNotContainPlayer  = fmt.Errorf("user is not in room")
	ErrNotYourTurn               = fmt.Errorf("not your turn")
	ErrInvalidMove               = fmt.Errorf("invalid move")
)
