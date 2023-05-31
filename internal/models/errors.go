package models

import (
	"github.com/pkg/errors"
)

var (
	ErrUnauthorized      = errors.New("unauthorized")
	ErrForbidden         = errors.New("forbidden")
	ErrBadRequest        = errors.New("bad request")
	ErrNotFound          = errors.New("not found")
	ErrConflict          = errors.New("conflict")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)

var (
	ErrGameCannotBeStarted = errors.New("game cannot be started")
	ErrGameNotInProgress   = errors.New("game not in progress")
	ErrPlayerNotInGame     = errors.New("player not in game")
	ErrNotYourTurn         = errors.New("not your turn")
	ErrInvalidMove         = errors.New("invalid move")
)
