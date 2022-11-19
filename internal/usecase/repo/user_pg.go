package repo

import (
	"auth_reg/internal/entity"
	"auth_reg/internal/usecase"
	"auth_reg/pkg/postgres"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type UserRepo struct {
	pg *postgres.Postgres
}

var _ usecase.UserRp = (*UserRepo)(nil)

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg: pg}
}

func (u *UserRepo) StoreUser(ctx context.Context, user entity.User) error {
	query := `insert into "user" (id, email, phone, password, status, secret_question_id, secret_question_answer) values ($1, $2, $3, $4, $5, $6, $7) returning id`

	rows, err := u.pg.Pool.Query(ctx, query, uuid.New(), user.Email, user.Phone, user.Password, "pending", user.SecretQuestionID, user.SecretQuestionAnswer)
	if err != nil {
		return fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return fmt.Errorf("error in creating user")
	}
	return nil
}

func (u *UserRepo) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	query := `select * from "user" where email=$1`

	rows, err := u.pg.Pool.Query(ctx, query, email)
	if err != nil {
		return entity.User{}, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var user entity.User
	if rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Phone,
			&user.Password,
			&user.Status,
			&user.SecretQuestionID,
			&user.SecretQuestionAnswer,
		)
		if err != nil {
			return entity.User{}, fmt.Errorf("error parsing user by email: %w", err)
		}
	}
	return user, nil
}

func (u *UserRepo) CheckExistenceByEmail(ctx context.Context, email string) (bool, error) {
	query := `select exists (select * from "user" where email=$1)`

	rows, err := u.pg.Pool.Query(ctx, query, email)
	if err != nil {
		return false, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var exists bool
	if rows.Next() {
		err = rows.Scan(&exists)
		if err != nil {
			return false, fmt.Errorf("error in parsing result exists user by email: %w", err)
		}
	}
	return exists, nil
}
