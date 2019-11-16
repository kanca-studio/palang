package user

import (
	"kanca-studio/palang/service/base"
)

type IdentifierType int

const (
	Email IdentifierType = iota
	PhoneNumber
	Username
)

type Service interface {
	base.InterfaceBaseRepository
	CreateUser(identifierType IdentifierType, identifier, password string) (Model, error)
	GetUserByIdentifier(identifierType IdentifierType, identifier string) (Model, error)
}

func NewService(repo repository) Service {
	s := service{}
	s.BaseService.Repository = &repo
	s.repo = repo

	return &s
}

type service struct {
	base.BaseService
	repo repository
}

func (s *service) CreateUser(identifierType IdentifierType, identifier, hashPassword string) (Model, error) {

	var data Model
	switch identifierType {
	case Email:
		data = Model{
			Email:    identifier,
			Password: hashPassword,
			Verified: true,
		}
	case PhoneNumber:
		data = Model{
			PhoneNumber: identifier,
			Password:    hashPassword,
			Verified:    true,
		}
	case Username:
		data = Model{
			Username: identifier,
			Password: hashPassword,
			Verified: true,
		}
	}

	err := s.Create(data)
	return data, err
}

func (s *service) GetUserByIdentifier(identifierType IdentifierType, identifier string) (Model, error) {
	var err error
	var result interface{}

	switch identifierType {
	case Email:
		result, err = s.Find(Model{Email: identifier})
	case PhoneNumber:
		result, err = s.Find(Model{PhoneNumber: identifier})
	case Username:
		result, err = s.Find(Model{Username: identifier})
	}
	if err != nil {
		return Model{}, err
	}
	user := result.(Model)
	return user, err
}
