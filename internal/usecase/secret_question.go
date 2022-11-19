package usecase

import (
	"auth_reg/internal/entity"
	"context"
)

type SecretQuestionUseCase struct {
	repo SecretQuestionRp
}

var _ SecretQuestionContract = (*SecretQuestionUseCase)(nil)

func NewSecretQuestionUseCase(repo SecretQuestionRp) *SecretQuestionUseCase {
	return &SecretQuestionUseCase{repo: repo}
}

func (s *SecretQuestionUseCase) GetAllQuestions(ctx context.Context) ([]entity.SecretQuestion, error) {
	return s.repo.GetAllQuestions(ctx)
}
