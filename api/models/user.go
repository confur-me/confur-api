package models

import (
	"github.com/confur-me/confur-api/db"
	"time"
)

type User struct {
	ID                   uint   `gorm:"primary_key"`
	Email                string `sql:"index" binding:"required"`
	Name                 string
	Password             string
	PasswordConfirmation string
	CreatedAt            time.Time
	ConfirmedAt          time.Time
	ConfirmationToken    string `sql:"type:text"`
	SignInToken          string `sql:"type:text"`
	ResetPasswordToken   string `sql:"type:text"`
}

func UserById(id string) User {
	var resource User
	d, err := db.Connection()
	if err == nil {
		d.Where("id = ?", id).First(&resource)
	}
	return resource
}
