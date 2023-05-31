package postgres

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"

	"go-labs-game-platform/internal/ctxpkg"
)

func (r *Repo) Begin(ctx context.Context) (context.Context, error) {
	tx := ctxpkg.GetDBTransaction(ctx)
	if tx != nil {
	}

	dbTx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	return ctxpkg.SetDBTransaction(ctx, &ctxpkg.DBTransaction{Tx: dbTx}), nil
}

func (r *Repo) Commit(ctx context.Context) error {
	tx := ctxpkg.GetDBTransaction(ctx)
	if tx == nil {
		return nil
	}

	if err := tx.Tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Repo) Rollback(ctx context.Context) {
	tx := ctxpkg.GetDBTransaction(ctx)
	if tx == nil {
		return
	}

	if err := tx.Tx.Rollback(); err != nil {
		logrus.Errorf("Rollback database transaction: %s", err)
		return
	}
}

func (r *Repo) getQuery(ctx context.Context, model interface{}) *pg.Query {
	tx := ctxpkg.GetDBTransaction(ctx)
	if tx != nil {
		return tx.Tx.ModelContext(ctx, model)
	}

	return r.db.ModelContext(ctx, model)
}
