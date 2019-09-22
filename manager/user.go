package manager

import (
	"kanca-studio/palang/service/email"
	"kanca-studio/palang/service/user"
)

type userManager struct {
	userService  user.Service
	emailService email.Service
}

func (m *userManager) Register(email, password string) error {
	data, err := m.userService.Register(email, password)
	if err != nil {
		return err
	}
	token := "" //TODO generated
	m.emailService.SendActivationToken(data.Email, token)
	return nil
}
