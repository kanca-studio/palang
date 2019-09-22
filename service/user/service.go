package user

import (
	"kanca-studio/palang/service/base"
)

func NewService(repository Repository) Service {
	s := Service{}
	s.BaseService.Repository = &repository
	s.repo = repository

	return s
}

type Service struct {
	base.BaseService
	repo Repository
}