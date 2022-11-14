package v1

import (
	"auth_reg/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine,
	u usecase.UserContract,
	r usecase.RegisterContract,
	j usecase.JwtContract) {

	h := handler.Group("/v1")

	{
		newRegisterRoutes(h, r, j)
		newLoginRoutes(h, u, j)
		newInfoRoutes(h, j)
		newLogoutRoutes(h, j)
	}
}
