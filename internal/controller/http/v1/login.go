package v1

import (
	"auth_reg/internal/entity"
	"auth_reg/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginRoutes struct {
	u usecase.UserContract
	j usecase.JwtContract
}

func newLoginRoutes(handler *gin.RouterGroup, u usecase.UserContract, j usecase.JwtContract) {
	lg := loginRoutes{u: u, j: j}

	handler.POST("/login", lg.login)
}

// @Summary Login
// @Tags login
// @Description login
// @ID login-user
// @Accept json
// @Produce json
// @Param input body loginRequestDTO true "Enter login, password, type"
// @Success 200 {object} nil
// @Failure 400 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /v1/login [post]
func (l *loginRoutes) login(c *gin.Context) {
	var lg loginRequestDTO
	if err := c.ShouldBindJSON(&lg); err != nil {
		errorResponse(c, http.StatusBadRequest, "cannot parse user creds")
		return
	}
	err := validateLogin(lg)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if lg.Type == "email" {
		err = l.j.CompareUserPasswordByEmail(c.Request.Context(), entity.User{
			Email:    lg.Login,
			Password: lg.Password,
		})
	} else {
		err = l.j.CompareUserPasswordByPhone(c.Request.Context(), entity.User{
			Phone:    lg.Login,
			Password: lg.Password,
		})
	}
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	email := lg.Login
	if lg.Type == "phone" {
		user, err := l.u.GetUserByPhone(c.Request.Context(), lg.Login)
		if err != nil {
			errorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		email = user.Email
	}
	token, err := l.j.GenerateToken(email)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "cannot generate token")
		return
	}
	c.Header("Authorization", "Bearer: "+token)
	c.JSON(http.StatusOK, nil)
}
