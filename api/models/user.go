package models

import (
	"github.com/confur-me/confur-api/db"
	"time"
)

type User struct {
	ID                   uint   `gorm:"primary_key"`
	Email                string `sql:"index" binding:"required"`
	Name                 string
	Password             string `json:"-"`
	PasswordConfirmation string `sql:"-" json:"-"`
	ConfirmationToken    string `sql:"type:text" json:"-"`
	SignInToken          string `sql:"type:text" json:"-"`
	ResetPasswordToken   string `sql:"type:text" json:"-"`
	CreatedAt            time.Time
	ConfirmedAt          time.Time
}

func UserById(id string) User {
	var resource User
	if d, ok := db.Connection(); ok {
		d.Where("id = ?", id).First(&resource)
	}
	return resource
}
