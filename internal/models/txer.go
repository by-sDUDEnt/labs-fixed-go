package models

import "context"

type TXer interface {
	Begin(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) error
}
