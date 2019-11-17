package schema

type BaseResponse struct {
	Message string `json:"message" `
}

type ErrorResponse struct {
	Message string `json:"message" `
	Error   string `json:"error" `
}

type ResponseRegister struct {
	BaseResponse
}

type ResponseLogin struct {
	BaseResponse
	Token string `json:"token" `
}

type ResponseMe struct {
	BaseResponse
	Data struct {
		Name        string `json:"name"`
		Email       string `json:"email,omitempty"`
		PhoneNumber string `json:"phone_number,omitempty"`
		Username    string `json:"username,omitempty"`
	} `json:"data"`
}
