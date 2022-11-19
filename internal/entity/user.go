package entity

import "github.com/google/uuid"

type User struct {
	ID                   uuid.UUID `json:"id"`
	Email                string    `json:"email"`
	Phone                string    `json:"phone"`
	Password             string    `json:"password"`
	Status               string    `json:"status"`
	SecretQuestionID     uuid.UUID `json:"secret_question_id"`
	SecretQuestionAnswer string    `json:"secret_question_answer"`
}
