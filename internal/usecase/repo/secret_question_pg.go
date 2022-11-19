package repo

import (
	"auth_reg/internal/entity"
	"auth_reg/internal/usecase"
	"auth_reg/pkg/postgres"
	"context"
	"fmt"
)

type SecretQuestionRepo struct {
	pg *postgres.Postgres
}

func NewSecretQuestionRepo(pg *postgres.Postgres) *SecretQuestionRepo {
	return &SecretQuestionRepo{pg: pg}
}

var _ usecase.SecretQuestionRp = (*SecretQuestionRepo)(nil)

func (s *SecretQuestionRepo) GetAllQuestions(ctx context.Context) ([]entity.SecretQuestion, error) {
	query := `select * from secret_question`

	rows, err := s.pg.Pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var questions []entity.SecretQuestion
	for rows.Next() {
		var question entity.SecretQuestion
		err = rows.Scan(&question.ID, &question.Question)
		if err != nil {
			return nil, fmt.Errorf("error in parsing question: %w", err)
		}
		questions = append(questions, question)
	}
	return questions, nil
}
