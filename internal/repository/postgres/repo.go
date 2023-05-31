package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/extra/pgdebug"
	"github.com/go-pg/pg/v10"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"

	"go-labs-game-platform/internal/config"
)

type Repo struct {
	db *pg.DB
}

func New() (*Repo, error) {
	cfg := config.Get().DB

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Database,
	})

	db.AddQueryHook(pgdebug.DebugHook{
		Verbose: config.Get().DB.Debug,
	})

	if err := db.Ping(context.Background()); err != nil {
		return nil, err
	}

	logrus.Info("Postgres: OK")

	return &Repo{db: db}, nil
}

func (r *Repo) Close() error {
	return r.db.Close()
}
