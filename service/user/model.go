package user

import "github.com/jinzhu/gorm"

type Model struct {
	gorm.Model
	Name               string `gorm:"column:name"`
	Email              string `gorm:"column:email"`
	Password           string `gorm:"column:password"`
	Verified           bool   `gorm:"column:verified"`
	VerificationToken  string `gorm:"column:verification_token"`
}
