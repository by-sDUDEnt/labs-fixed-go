package ctxpkg

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"

	"go-labs-game-platform/internal/models"
)

type ctxKey string

var (
	userIDKey     = ctxKey("user_id")
	userKey       = ctxKey("users")
	dbTransaction = ctxKey("db_transaction")
)

type DBTransaction struct {
	Tx *pg.Tx
}

func SetDBTransaction(ctx context.Context, tx *DBTransaction) context.Context {
	return context.WithValue(ctx, dbTransaction, tx)
}

func GetDBTransaction(ctx context.Context) *DBTransaction {
	tx, _ := ctx.Value(dbTransaction).(*DBTransaction)
	return tx
}

func SetUserID(ctx context.Context, id uuid.UUID) context.Context {
	return context.WithValue(ctx, userIDKey, id)
}

func GetUserID(ctx context.Context) uuid.UUID {
	userID, _ := ctx.Value(userIDKey).(uuid.UUID)
	return userID
}

func SetUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func GetUser(ctx context.Context) *models.User {
	user, _ := ctx.Value(userKey).(*models.User)
	return user
}
