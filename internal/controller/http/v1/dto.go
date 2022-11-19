package v1

type registerRequestDTO struct {
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	Password             string `json:"password"`
	SecretQuestionID     string `json:"secretQuestionID"`
	SecretQuestionAnswer string `json:"secretQuestionAnswer"`
}
