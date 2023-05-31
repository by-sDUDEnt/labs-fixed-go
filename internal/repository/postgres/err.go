package postgres

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"

	"go-labs-game-platform/internal/models"
)

func (r *Repo) err(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, pg.ErrNoRows) {
		return models.ErrNotFound
	}

	pgErr, ok := err.(pg.Error)
	if ok && pgErr.IntegrityViolation() {
		return fmt.Errorf("%w: %s", models.ErrConflict, err)
	}

	return err
}
