package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" pg:",pk,type:uuid,default:uuid_generate_v4()"`
	Username  string    `json:"username" pg:",unique,notnull"`
	CreatedAt time.Time `json:"-" pg:",notnull,default:now()"`
}
