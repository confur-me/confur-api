package models

import (
	"github.com/confur-me/confur-api/db"
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Email     *string   `sql:"not null;index" binding:"required" json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
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
