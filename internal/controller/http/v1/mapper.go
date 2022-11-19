package v1

import (
	"auth_reg/internal/entity"
	"fmt"
	"github.com/google/uuid"
)

func userRegisterToEntity(dto registerRequestDTO) (entity.User, error) {
	secretQuestionID, err := uuid.Parse(dto.SecretQuestionID)
	if err != nil {
		return entity.User{}, fmt.Errorf("invalid secretQuestionID uuid: %w", err)
	}
	return entity.User{
		Email:                dto.Email,
		Phone:                dto.Phone,
		Password:             dto.Password,
		SecretQuestionID:     secretQuestionID,
		SecretQuestionAnswer: dto.SecretQuestionAnswer,
	}, nil
}
