package manager

import (
	"errors"

	"github.com/kanca-studio/palang/service/auth"
	"github.com/kanca-studio/palang/service/user"
)

type User struct {
	userService user.Service
	authService auth.Service
}

func NewUser(userService user.Service, authService auth.Service) User {
	return User{
		userService: userService,
		authService: authService,
	}
}

func (m *User) Register(identifierTypeStr, identifier, password string) error {

	identifierType := m.userService.IdentifierTypeToConst(identifierTypeStr)
	hash, _ := m.authService.HashPassword(password)
	_, err := m.userService.CreateUser(identifierType, identifier, hash)
	if err != nil {
		return err
	}

	//	token := "" //TODO generated
	//	m.emailService.SendActivationToken(data.Email, token)
	return nil
}

func (m *User) Login(identifierTypeStr, identifier, password string) (string, error) {

	identifierType := m.userService.IdentifierTypeToConst(identifierTypeStr)
	data, err := m.userService.GetUserByIdentifier(identifierType, identifier)

	if err != nil {
		return "", err
	}

	if !m.authService.CheckPasswordHash(password, data.Password) {
		return "", errors.New("please check again username or password")
	}

	if !data.Verified {
		return "", errors.New("User not verified")
	}

	token, err := m.authService.CreateToken(data.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m *User) ValidateToken(token string) error {
	if _, err := m.authService.ValidateToken(token); err != nil {
		return err
	}
	return nil
}

func (m *User) Me(token string) (user.Model, error) {
	claim, err := m.authService.ValidateToken(token)

	if err != nil {
		return user.Model{}, err
	}
	var dataUser user.Model
	id := uint(claim["sub"].(float64))
	if err := m.userService.FindById(id, &dataUser); err != nil {
		return user.Model{}, err
	}

	return dataUser, err
}
