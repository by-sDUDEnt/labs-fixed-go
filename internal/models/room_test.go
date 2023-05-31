//package models
//
//import (
//	"github.com/google/uuid"
//	"testing"
//)
//
//func TestRoom_IsFull(t *testing.T) {
//	type fields struct {
//		ID      uuid.UUID
//		Players []uuid.UUID
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   bool
//	}{
//		{
//			name: "room is full",
//			fields: fields{
//				ID:      uuid.New(),
//				Players: []uuid.UUID{uuid.New(), uuid.New()},
//			},
//			want: true,
//		},
//		{
//			name: "room is not full",
//			fields: fields{
//				ID:      uuid.New(),
//				Players: []uuid.UUID{uuid.New()},
//			},
//			want: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//
//			r := &Room{
//				ID:        tt.fields.ID,
//				PlayerIDs: tt.fields.Players,
//			}
//
//			if got := r.IsFull(); got != tt.want {
//				t.Errorf("Room.IsFull() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestRoom_ContainsPlayer(t *testing.T) {
//	type fields struct {
//		ID      uuid.UUID
//		Players []uuid.UUID
//	}
//	type args struct {
//		userID uuid.UUID
//	}
//	userID := uuid.New()
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   bool
//	}{
//		{
//			"room contains player",
//			fields{
//				ID:      uuid.New(),
//				Players: []uuid.UUID{userID},
//			},
//			args{
//				userID: userID,
//			},
//			true,
//		},
//		{
//			"room does not contain player",
//			fields{
//				ID:      uuid.New(),
//				Players: []uuid.UUID{uuid.New()},
//			},
//			args{
//				userID: uuid.New(),
//			},
//			false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//
//			r := &Room{
//				ID:        tt.fields.ID,
//				PlayerIDs: tt.fields.Players,
//			}
//
//			if got := r.ContainsPlayer(tt.args.userID); got != tt.want {
//				t.Errorf("Room.ContainsPlayer() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

package models

import (
	"testing"

	"github.com/google/uuid"
)

func TestRoom_AddPlayer(t *testing.T) {
	room := Room{}
	playerID := uuid.New()
	room.AddPlayer(playerID)

	if len(room.PlayerIDs) != 1 {
		t.Errorf("Expected player to be added to the room")
	}

	if room.PlayerIDs[0] != playerID {
		t.Errorf("Expected player ID to match the added player")
	}
}

func TestRoom_RemovePlayer(t *testing.T) {
	room := Room{}
	playerID := uuid.New()
	room.PlayerIDs = append(room.PlayerIDs, playerID)

	removed := room.RemovePlayer(playerID)

	if !removed {
		t.Errorf("Expected player to be removed from the room")
	}

	if len(room.PlayerIDs) != 0 {
		t.Errorf("Expected player to be removed from the room")
	}

	// Try removing a player that doesn't exist
	removed = room.RemovePlayer(uuid.New())

	if removed {
		t.Errorf("Expected player to not be removed from the room")
	}
}

func TestRoom_ContainsPlayer(t *testing.T) {
	room := Room{}
	playerID := uuid.New()
	room.PlayerIDs = append(room.PlayerIDs, playerID)

	contains := room.ContainsPlayer(playerID)

	if !contains {
		t.Errorf("Expected room to contain the player")
	}

	contains = room.ContainsPlayer(uuid.New())

	if contains {
		t.Errorf("Expected room to not contain the player")
	}
}

func TestRoom_IsEmpty(t *testing.T) {
	room := Room{}

	isEmpty := room.IsEmpty()

	if !isEmpty {
		t.Errorf("Expected room to be empty")
	}

	room.PlayerIDs = append(room.PlayerIDs, uuid.New())

	isEmpty = room.IsEmpty()

	if isEmpty {
		t.Errorf("Expected room to not be empty")
	}
}

func TestRoom_IsFull(t *testing.T) {
	room := Room{}

	isFull := room.IsFull()

	if isFull {
		t.Errorf("Expected room to not be full")
	}

	room.PlayerIDs = append(room.PlayerIDs, uuid.New())
	room.PlayerIDs = append(room.PlayerIDs, uuid.New())

	isFull = room.IsFull()

	if !isFull {
		t.Errorf("Expected room to be full")
	}
}

func TestRoom_CanBeStarted(t *testing.T) {
	room := Room{}

	canBeStarted := room.CanBeStarted()

	if canBeStarted {
		t.Errorf("Expected room to not be able to be started")
	}

	room.PlayerIDs = append(room.PlayerIDs, uuid.New())
	room.PlayerIDs = append(room.PlayerIDs, uuid.New())

	canBeStarted = room.CanBeStarted()

	if !canBeStarted {
		t.Errorf("Expected room to be able to be started")
	}
}

func TestRoom_IndexOfPlayer(t *testing.T) {
	room := Room{}
	player1 := uuid.New()
	player2 := uuid.New()
	room.PlayerIDs = append(room.PlayerIDs, player1)
	room.PlayerIDs = append(room.PlayerIDs, player2)

	index := room.IndexOfPlayer(player1)

	if index != 0 {
		t.Errorf("Expected index to be 0 for player1")
	}

	index = room.IndexOfPlayer(player2)

	if index != 1 {
		t.Errorf("Expected index to be 1 for player2")
	}

	index = room.IndexOfPlayer(uuid.New())

	if index != -1 {
		t.Errorf("Expected index to be -1 for player that doesn't exist")
	}
}

func TestRoom_HasWinner(t *testing.T) {
	// Create a room with a winning state
	room := Room{}
	room.Table[0][0] = 1
	room.Table[0][1] = 1
	room.Table[0][2] = 1
	room.Table[0][3] = 1

	hasWinner := room.HasWinner()

	if !hasWinner {
		t.Errorf("Expected room to have a winner")
	}

	// Create a room without a winning state
	room = Room{}
	room.Table[0][0] = 1
	room.Table[0][1] = 1
	room.Table[0][2] = 1

	hasWinner = room.HasWinner()

	if hasWinner {
		t.Errorf("Expected room to not have a winner")
	}
}
