package user

import (
	"github.com/jinzhu/gorm"
	"kanca-studio/palang/service/base"
)

func NewRepository(db *gorm.DB) repository {
	return repository{
		base.NewBaseRepository(db, Model{}),
	}
}

type repository struct {
	base.BaseRepository
}

func (*repository) anotherFunction() bool {
	return true
}
