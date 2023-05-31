package tictactoe

type Side string

const (
	SideX Side = "x"
	SideO Side = "o"
)

type (
	GameStarted struct {
		YourSide Side
	}

	OpponentMoved struct {
		Position int
	}

	YourTurn struct {
	}

	Move struct {
		Position int
	}
)
