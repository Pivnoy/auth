package v1

import (
	"auth_reg/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type secretQuestionRoutes struct {
	su usecase.SecretQuestionContract
}

func newSecretQuestionRoutes(handler *gin.RouterGroup, su usecase.SecretQuestionContract) {
	sr := secretQuestionRoutes{su: su}

	handler.GET("/questions", sr.getSecretQuestions)
}

// @Summary GetQuestions
// @Tags questions
// @Description Get all secret questions
// @ID get-question
// @Accept json
// @Produce json
// @Success 200 {object} secretQuestionsResponseDTO
// @Failure 500 {object} errResponse
// @Router /v1/questions [get]
func (s *secretQuestionRoutes) getSecretQuestions(c *gin.Context) {
	listQuestions, err := s.su.GetAllQuestions(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, secretQuestionsResponseDTO{listQuestions})
}
