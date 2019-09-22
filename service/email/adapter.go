package email

type interfaceAdapter interface {
	sendEmail(param interface{}) interface{}
}
