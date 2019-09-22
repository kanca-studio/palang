package email

func NewSendgridAdapater() interfaceAdapter {
	return &sendgrid{}
}

type sendgrid struct {
}

func (s *sendgrid) sendEmail(param interface{}) interface{} {
	return nil
}
