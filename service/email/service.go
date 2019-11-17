package email

type Service interface {
	SendActivationToken(email, token string) error
}

func NewService(adapter interfaceAdapter) Service {
	return &service{adapter: adapter}
}

type service struct {
	adapter interfaceAdapter
}

func (s *service) SendActivationToken(email, token string) error {

	//TODO s.adapter.sendEmail()
	return nil
}
