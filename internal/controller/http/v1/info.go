package v1

import (
	"auth_reg/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type infoRoutes struct {
	j usecase.JwtContract
}

func newInfoRoutes(handler *gin.RouterGroup, j usecase.JwtContract) {
	i := infoRoutes{j: j}

	handler.GET("/info", i.info)
}

type infoResponse struct {
	Email string `json:"email"`
}

func (i *infoRoutes) info(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		errorResponse(c, http.StatusUnauthorized, "error in header format")
	}
	headerParts := strings.Split(auth, " ")
	if len(headerParts) != 2 {
		errorResponse(c, http.StatusUnauthorized, "cannot parse token")
		return
	}
	if headerParts[0] != "Bearer" {
		errorResponse(c, http.StatusUnauthorized, "cannot find Bearer")
		return
	}
	token, err := i.j.CheckToken(headerParts[1])
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "cannot check token")
		return
	}
	c.JSON(http.StatusOK, infoResponse{token})
}
