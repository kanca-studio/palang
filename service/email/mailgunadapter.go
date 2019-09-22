package email

func NewMailgunAdapater() interfaceAdapter {
	return &mailgun{}
}

type mailgun struct {
}

func (m *mailgun) sendEmail(param interface{}) interface{} {
	return nil
}
