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

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *registerRoutes) register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, "cannot parse request credentials")
		return
	}

	err := r.rg.CreateNewUser(c.Request.Context(), req.Email, req.Password)
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
