package models

import (
	"github.com/confur-me/confur-api/db"
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Email     string `sql:"index" binding:"required"`
	Name      string
	CreatedAt time.Time
}

type userService struct {
	Service
}

func NewUserService(params map[string]interface{}) *userService {
	s := new(userService)
	s.params = params
	return s
}

func (this *userService) User(id string) (*User, error) {
	var (
		resource User
		err      error
	)
	if conn, ok := db.Connection(); ok {
		err = conn.Where("id = ?", id).First(&resource).Error
	}
	return &resource, err
}
