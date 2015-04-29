package models

import (
	"github.com/confur-me/confur-api/db"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email                string `sql:"index"`
	Name                 string
	Password             string
	PasswordConfirmation string
	ConfirmedAt          *time.Time
}

func UserById(id string) User {
	var resource User
	d, err := db.Connection()
	if err == nil {
		d.Where("id = ?", id).First(&resource)
	}
	return resource
}
