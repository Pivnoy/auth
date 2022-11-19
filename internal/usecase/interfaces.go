package usecase

import (
	"auth_reg/internal/entity"
	"context"
)

type (
	UserRp interface {
		StoreUser(context.Context, entity.User) error
		GetUserByEmail(context.Context, string) (entity.User, error)
		CheckExistenceByEmail(context.Context, string) (bool, error)
	}

	UserContract interface {
		StoreUser(context.Context, entity.User) error
		GetUserByEmail(context.Context, string) (entity.User, error)
		CheckExistenceByEmail(context.Context, string) (bool, error)
	}

	JwtContract interface {
		CompareUserPassword(context.Context, entity.User) error
		GenerateToken(string) (string, error)
		CheckToken(string) (string, error)
	}

	RegisterContract interface {
		CreateNewUser(context.Context, entity.User) error
	}

	SecretQuestionContract interface {
		GetAllQuestions(context.Context) ([]entity.SecretQuestion, error)
	}

	SecretQuestionRp interface {
		GetAllQuestions(context.Context) ([]entity.SecretQuestion, error)
	}
)
