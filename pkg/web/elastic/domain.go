package elastic

import (
	"github.com/google/uuid"
	"time"
)

type Document struct {
	ID        uuid.UUID `json:"id" elk:"id"`
	CreatedAt time.Time `json:"created_at" elk:"created_at"`
}

func NewDocument() Document {
	return Document{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
	}
}
