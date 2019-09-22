package user

import (
	"kanca-studio/palang/service/base"
)

func NewService(repo repository) service {
	s := service{}
	s.BaseService.Repository = &repo
	s.repo = repo

	return s
}

type service struct {
	base.BaseService
	repo repository
}
