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

	handler.POST("/auth/register", r.register)
}

// @Summary Register
// @Tags auth
// @Description Register new user
// @ID register-user
// @Accept json
// @Produce json
// @Param input body registerRequestDTO true "Enter info user"
// @Success 200 {object} nil
// @Failure 400 {object} errResponse
// @Failure 401 {object} errResponse
// @Router /v1/auth/register [post]
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
	c.JSON(http.StatusOK, nil)
}
