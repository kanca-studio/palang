package user

import "github.com/jinzhu/gorm"

type Model struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);column:name"`
	Email    string `gorm:"type:varchar(255);column:email"`
	Password string `gorm:"type:varchar(255);column:password"`
}