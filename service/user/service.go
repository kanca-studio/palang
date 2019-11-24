package user

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/kanca-studio/palang/service/base"
)

type IdentifierType int

const (
	Email IdentifierType = iota
	PhoneNumber
	Username
)

type Service interface {
	base.InterfaceBaseRepository
	CreateUser(identifierTypeStr IdentifierType, identifier, password string) (Model, error)
	GetUserByIdentifier(identifierTypeStr IdentifierType, identifier string) (Model, error)
	IdentifierTypeToConst(identifierType string) IdentifierType
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

	//check user exist
	checking, err := s.GetUserByIdentifier(identifierType, identifier)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return data, err
	}
	if checking.ID != 0 {
		return data, errors.New("user alerady exist")
	}

	err = s.Create(&data)
	return data, err
}

func (s *service) GetUserByIdentifier(identifierType IdentifierType, identifier string) (Model, error) {
	var err error

	var result Model

	switch identifierType {
	case Email:
		err = s.Find(Model{Email: identifier}, &result)
	case PhoneNumber:
		err = s.Find(Model{PhoneNumber: identifier}, &result)
	case Username:
		err = s.Find(Model{Username: identifier}, &result)
	}
	if err != nil {
		return Model{}, err
	}

	return result, err
}

func (s *service) IdentifierTypeToConst(identifierType string) IdentifierType {
	if identifierType == "Username" {
		return Username
	}
	if identifierType == "PhoneNumber" {
		return PhoneNumber
	}
	return Email
}
