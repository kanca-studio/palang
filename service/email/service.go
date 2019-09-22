package email

func NewService(adapter interfaceAdapter) service {
	return service{adapter: adapter}
}

type service struct {
	adapter interfaceAdapter
}

func (s *service) sendActiviationToken(token string) error {

	//TODO s.adapter.sendEmail()
	return nil
}
