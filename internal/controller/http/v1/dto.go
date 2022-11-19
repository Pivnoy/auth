package v1

type registerRequestDTO struct {
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	Password             string `json:"password"`
	SecretQuestionID     string `json:"secretQuestionID"`
	SecretQuestionAnswer string `json:"secretQuestionAnswer"`
}

type secretQuestionsResponseDTO struct {
	Questions []secretQuestionResponseDTO `json:"data"`
}

type loginRequestDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

// data value label

type secretQuestionResponseDTO struct {
	Value string `json:"value"`
	Label string `json:"label"`
}
