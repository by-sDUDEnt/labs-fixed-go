package models

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
