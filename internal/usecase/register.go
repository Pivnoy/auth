package usecase

import (
	"auth_reg/internal/entity"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUseCase struct {
	u UserContract
}

var _ RegisterContract = (*RegisterUseCase)(nil)

func NewRegisterUseCase(u UserContract) *RegisterUseCase {
	return &RegisterUseCase{
		u: u,
	}
}

func (r *RegisterUseCase) CreateNewUser(ctx context.Context, user entity.User) error {
	exists, err := r.u.CheckExistenceByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("user already exists")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return fmt.Errorf("error in hashing psswrd: %v", err)
	}
	secretAnswerHash, err := bcrypt.GenerateFromPassword([]byte(user.SecretQuestionAnswer), 14)
	user.Password = string(passwordHash)
	user.SecretQuestionAnswer = string(secretAnswerHash)
	return r.u.StoreUser(ctx, user)
}
