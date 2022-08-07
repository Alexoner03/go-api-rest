package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
}
