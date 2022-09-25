package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	"jwt/helpers"
)

type User struct {
	GormModel
	FullName string
	Email    string
	Password string
	Products []Product
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
