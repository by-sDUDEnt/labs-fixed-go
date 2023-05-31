package models

import (
	"time"

	"github.com/google/uuid"
)

const MaxPlayers = 2

type RoomList []Room

type Room struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	GameType GameType `json:"game_type" pg:",notnull"`

	Status GameStatus `json:"status" pg:",notnull"`

	PlayerIDs           []uuid.UUID `json:"players"`
	CurrentMovePlayerID uuid.UUID   `json:"current_move_player_id"`
	NextMovePlayerID    uuid.UUID   `json:"next_move_player_id"`
	Table               [6][7]int   `json:"table"`
}

func (r *Room) AddPlayer(player uuid.UUID) {
	r.PlayerIDs = append(r.PlayerIDs, player)
}

func (r *Room) RemovePlayer(player uuid.UUID) bool {
	for i, p := range r.PlayerIDs {
		if p == player {
			r.PlayerIDs = append(r.PlayerIDs[:i], r.PlayerIDs[i+1:]...)
			return true
		}
	}

	return false
}

func (r *Room) ContainsPlayer(userID uuid.UUID) bool {
	for _, player := range r.PlayerIDs {
		if player == userID {
			return true
		}
	}

	return false
}

func (r *Room) IsEmpty() bool {
	return len(r.PlayerIDs) == 0
}

func (r *Room) IsFull() bool {
	return len(r.PlayerIDs) == MaxPlayers
}

func (r *Room) CanBeStarted() bool {
	return r.IsFull()
}

func (r *Room) IndexOfPlayer(id uuid.UUID) int {
	for i, player := range r.PlayerIDs {
		if player == id {
			return i
		}
	}

	return -1
}

func (r *Room) HasWinner() bool {
	for row := 0; row < 6; row++ {
		for col := 0; col < 7; col++ {
			if r.Table[row][col] == 0 {
				continue
			}

			if col+3 < 7 &&
				r.Table[row][col] == r.Table[row][col+1] &&
				r.Table[row][col] == r.Table[row][col+2] &&
				r.Table[row][col] == r.Table[row][col+3] {
				return true
			}

			if row+3 < 6 {
				if r.Table[row][col] == r.Table[row+1][col] &&
					r.Table[row][col] == r.Table[row+2][col] &&
					r.Table[row][col] == r.Table[row+3][col] {
					return true
				}

				if col+3 < 7 &&
					r.Table[row][col] == r.Table[row+1][col+1] &&
					r.Table[row][col] == r.Table[row+2][col+2] &&
					r.Table[row][col] == r.Table[row+3][col+3] {
					return true
				}

				if col-3 >= 0 &&
					r.Table[row][col] == r.Table[row+1][col-1] &&
					r.Table[row][col] == r.Table[row+2][col-2] &&
					r.Table[row][col] == r.Table[row+3][col-3] {
					return true
				}
			}
		}
	}

	return false
}
