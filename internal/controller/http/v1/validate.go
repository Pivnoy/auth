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
