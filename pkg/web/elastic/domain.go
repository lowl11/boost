package elastic

import (
	"github.com/google/uuid"
	"time"
)

type Document struct {
	ID        uuid.UUID `json:"id" elk:"id"`
	CreatedAt time.Time `json:"created_at" elk:"created_at"`
}

func NewDocument(customID ...uuid.UUID) Document {
	document := Document{
		CreatedAt: time.Now(),
	}
	if len(customID) > 0 {
		document.ID = customID[0]
	} else {
		document.ID = uuid.New()
	}
	return document
}
