package user

import (
	"github.com/jinzhu/gorm"
	"github.com/kanca-studio/palang/service/base"
)

func WithGorm(db *gorm.DB) func(*repository) error {
	return func(r *repository) error {
		baserepo := base.NewBaseRepository(db, Model{})
		return r.setBaseRepo(&baserepo)
	}
}

func WithInterface(baserepo base.InterfaceBaseRepository) func(*repository) error {
	return func(r *repository) error {
		return r.setBaseRepo(baserepo)
	}
}

func NewRepository(options ...func(*repository) error) repository {
	repo := repository{}
	for _, option := range options {
		if err := option(&repo); err != nil {
			panic(err)
		}
	}
	return repo
}

type repository struct {
	base.InterfaceBaseRepository
}

func (r *repository) setBaseRepo(baserepo base.InterfaceBaseRepository) error {
	r.InterfaceBaseRepository = baserepo
	return nil
}
