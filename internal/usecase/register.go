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

func (r *RegisterUseCase) CreateNewUser(ctx context.Context, email string, password string) error {
	if len(email) < 4 || len(password) < 4 {
		return fmt.Errorf("invalid format of username or password")
	}
	exists, err := r.u.CheckExistenceByEmail(ctx, email)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("user already exists")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return fmt.Errorf("error in hashing psswrd: %v", err)
	}
	us := entity.User{
		Email:    email,
		Password: string(passwordHash),
	}
	return r.u.StoreUser(ctx, us)
}
