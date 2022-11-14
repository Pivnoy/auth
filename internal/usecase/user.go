package usecase

import (
	"auth_reg/internal/entity"
	"context"
)

type UserUseCase struct {
	repo UserRp
}

var _ UserContract = (*UserUseCase)(nil)

func NewUserUseCase(repo UserRp) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) StoreUser(ctx context.Context, user entity.User) error {
	return u.repo.StoreUser(ctx, user)
}

func (u *UserUseCase) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	return u.repo.GetUserByEmail(ctx, email)
}

func (u *UserUseCase) CheckExistenceByEmail(ctx context.Context, email string) (bool, error) {
	return u.repo.CheckExistenceByEmail(ctx, email)
}
