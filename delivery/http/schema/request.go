package schema

type ReqRegister struct {
	IdentifierType string `json:"identifierType" validate:"oneof=Email PhoneNumber Username"`
	Identifier     string `json:"identifier" validate:"required"`
	Password       string `json:"password" validate:"required"`
}

type ReqLogin struct {
	IdentifierType string `json:"identifierType" validate:"oneof=email username phonenumber"`
	Identifier     string `validate:"required"`
	Password       string `validate:"required"`
}
