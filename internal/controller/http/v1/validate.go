package v1

import (
	"auth_reg/internal/entity"
	"fmt"
)

func validateNewUser(user entity.User) (err error) {
	switch {
	case len(user.Email) < 4:
		err = fmt.Errorf("invalid format email")
	case len(user.Password) < 4:
		err = fmt.Errorf("invalid format password")
	case len(user.Phone) < 9 || len(user.Phone) > 12:
		err = fmt.Errorf("invalid format phone")
	}
	return err
}

func validateLogin(dto loginRequestDTO) (err error) {
	switch {
	case len(dto.Login) < 4:
		err = fmt.Errorf("invalid format login")
	case len(dto.Password) < 4:
		err = fmt.Errorf("invalid format password")
	case dto.Type != "email" && dto.Type != "phone":
		err = fmt.Errorf("invalid format type of login")
	}
	return err
}
