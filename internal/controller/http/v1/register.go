package v1

import (
	"auth_reg/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerRoutes struct {
	rg usecase.RegisterContract
	j  usecase.JwtContract
}

func newRegisterRoutes(handler *gin.RouterGroup, rg usecase.RegisterContract, j usecase.JwtContract) {
	r := registerRoutes{rg: rg, j: j}

	handler.POST("/register", r.register)
}

func (r *registerRoutes) register(c *gin.Context) {
	var req registerRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, "cannot parse request credentials")
		return
	}
	user, err := userRegisterToEntity(req)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = validateNewUser(user)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = r.rg.CreateNewUser(c.Request.Context(), user)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, "error in format user creds")
		return
	}
	token, err := r.j.GenerateToken(req.Email)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "cannot generate token")
		return
	}
	c.Header("Authorization", "Bearer: "+token)
	c.JSON(http.StatusOK, nil)
}
