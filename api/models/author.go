package models

import (
	"github.com/confur-me/confur-api/db"
	"time"
)

type Author struct {
	ID         uint `gorm:"primary_key"`
	FirstName  string
	LastName   string
	MiddleName string
	CreatedAt  time.Time
	Photo      string
	Events     []Event `gorm:"many2many:events_authors"`
}

func AuthorById(id string) Author {
	var resource Author
	d, err := db.Connection()
	if err == nil {
		d.Where("id = ?", id).First(&resource)
	}
	return resource
}
