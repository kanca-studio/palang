package user

import (
	"github.com/jinzhu/gorm"
	"kanca-studio/palang/service/base"
)

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		base.NewBaseRepository(db, Model{}),
	}
}

type Repository struct {
	base.BaseRepository
}

func (*Repository) anotherFunction() bool {
	return true
}