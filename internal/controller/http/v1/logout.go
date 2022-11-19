package v1

import (
	"auth_reg/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type logoutRoutes struct {
	j usecase.JwtContract
}

func newLogoutRoutes(handler *gin.RouterGroup, j usecase.JwtContract) {
	lg := logoutRoutes{j: j}

	handler.POST("/auth/logout", lg.logout)
}

// @Summary Logout
// @Tags auth
// @Description logout
// @ID logout-user
// @Accept json
// @Produce json
// @Success 200 {object} nil
// @Failure 400 {object} errResponse
// @Failure 401 {object} errResponse
// @Router /v1/auth/logout [post]
func (l *logoutRoutes) logout(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		errorResponse(c, http.StatusUnauthorized, "error in header format")
		return
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
	_, err := l.j.CheckToken(headerParts[1])
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "cannot check token")
		return
	}
	c.Header("Authorization", "")
	c.JSON(http.StatusOK, nil)
}
