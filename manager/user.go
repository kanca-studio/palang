package manager

import (
	"errors"
	"kanca-studio/palang/service/auth"
	"kanca-studio/palang/service/email"
	"kanca-studio/palang/service/user"
)

type userManager struct {
	userService  user.Service
	emailService email.Service
	authService  auth.Service
}

func (m *userManager) Register(identifierType user.IdentifierType, identifier, password string) error {

	hash, _ := m.authService.HashPassword(password)
	_, err := m.userService.CreateUser(identifierType, identifier, hash)
	if err != nil {
		return err
	}

	//	token := "" //TODO generated
	//	m.emailService.SendActivationToken(data.Email, token)
	return nil
}

func (m *userManager) Login(identifierType user.IdentifierType, identifier, password string) (string, error) {
	data, err := m.userService.GetUserByIdentifier(identifierType, identifier)

	if err != nil {
		return "", err
	}

	if !m.authService.CheckPasswordHash(password, data.Password) {
		return "", errors.New("please check again username or password")
	}

	if data.Verified == false {
		return "", errors.New("User not verified")
	}

	token, err := m.authService.CreateToken(data.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m *userManager) ValidateToken(token string) error {
	if _, err := m.authService.ValidateToken(token); err != nil {
		return err
	}
	return nil
}

func (m *userManager) Me(token string) (user.Model, error) {
	claim, err := m.authService.ValidateToken(token)

	if err != nil {
		return user.Model{}, err
	}

	id := claim["sub"].(uint)
	result, err := m.userService.FindById(id)
	if err != nil {
		return user.Model{}, err
	}
	dataUser := result.(user.Model)

	return dataUser, err
}
