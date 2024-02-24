package sql

import (
	"github.com/google/uuid"
	"time"
)

type Entity struct {
	ID        uuid.UUID `db:"id" ef:"pk"`
	CreatedAt time.Time `db:"created_at" ef:"default:now()"`
	UpdatedAt time.Time `db:"updated_at" ef:"default:now()"`
}
