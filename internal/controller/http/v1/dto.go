package v1

import "auth_reg/internal/entity"

type registerRequestDTO struct {
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	Password             string `json:"password"`
	SecretQuestionID     string `json:"secretQuestionID"`
	SecretQuestionAnswer string `json:"secretQuestionAnswer"`
}

type secretQuestionsResponseDTO struct {
	Questions []entity.SecretQuestion `json:"questions"`
}
