package entity

import "github.com/google/uuid"

type SecretQuestion struct {
	ID       uuid.UUID `json:"id"`
	Question string    `json:"question"`
}
