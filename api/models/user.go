package models

import (
	"github.com/confur-me/confur-api/db"
	"time"
)

type User struct {
	ID                   uint   `gorm:"primary_key"`
	Email                string `sql:"index"`
	Name                 string
	Password             string
	PasswordConfirmation string
	CreatedAt            time.Time
	ConfirmedAt          time.Time
}

func UserById(id string) User {
	var resource User
	d, err := db.Connection()
	if err == nil {
		d.Where("id = ?", id).First(&resource)
	}
	return resource
}
