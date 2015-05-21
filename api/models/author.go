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
	Events     []Event `gorm:"many2many:events_authors" json:",omitempty"`
}

type AuthorService struct {
	Service
}

func (this *AuthorService) Author() (*Author, error) {
	var (
		resource Author
		err      error
	)
	if conn, ok := db.Connection(); ok {
		if v, ok := this.params["id"]; ok {
			err = conn.Where("id = ?", v).First(&resource).Error
		}
	}
	return &resource, err
}
