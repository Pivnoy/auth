package v1

import (
	_ "auth_reg/docs"
	"auth_reg/internal/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(handler *gin.Engine,
	u usecase.UserContract,
	r usecase.RegisterContract,
	j usecase.JwtContract) {

	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	h := handler.Group("/v1")

	{
		newRegisterRoutes(h, r, j)
		newLoginRoutes(h, u, j)
		newInfoRoutes(h, j)
		newLogoutRoutes(h, j)
	}
}
