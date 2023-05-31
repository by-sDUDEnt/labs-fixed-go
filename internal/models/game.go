package models

import (
	"time"

	"github.com/google/uuid"
)

type GameType string
type GameStatus string

const (
	GameTypeFourInARow GameType = "four-in-a-row"
)

const (
	GameStatusWaitingForPlayers GameStatus = "waiting-for-players"
	GameStatusInProgress        GameStatus = "in-progress"
	GameStatusFinished          GameStatus = "finished"
)

type Game interface {
	baseGame() BaseGame

	CanBeStarted() bool

	Start() error

	Move(player Player, move any) error

	MaxPlayers() int
	MinPlayers() int
}

// BaseGame is the base game model. It contains all the fields that are common
// to all game. It must be embedded in all game models.
type BaseGame struct {
	ID        uuid.UUID `json:"id" pg:",pk,use_zero"`
	CreatedAt time.Time `json:"-" pg:",use_zero"`
	UpdatedAt time.Time `json:"-" pg:",use_zero"`

	Type   GameType   `json:"type" pg:",use_zero"`
	Status GameStatus `json:"status" pg:",use_zero"`

	Players []uuid.UUID `json:"players" pg:"-"`

	StartedAt  *time.Time `json:"started_at" pg:",use_zero"`
	FinishedAt *time.Time `json:"finished_at" pg:",use_zero"`
}

func NewGame(gameType GameType, players []uuid.UUID) Game {
	baseGame := BaseGame{
		Type:    gameType,
		Status:  GameStatusWaitingForPlayers,
		Players: players,
	}

	_ = baseGame

	switch gameType {
	default:
		return nil
	}
}

func (bg BaseGame) baseGame() BaseGame {
	return bg
}

func (bg BaseGame) Start() {
	now := time.Now()

	bg.Status = GameStatusInProgress
	bg.StartedAt = &now
}

func (bg BaseGame) Send(x Player, packet any) {
	panic("implement me")
}
