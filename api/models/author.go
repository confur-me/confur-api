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

type AuthorService struct {
	service
}

func (this *AuthorService) GetAuthor() (Author, bool) {
	var (
		resource Author
		success  bool
	)
	if d, err := db.Connection(); err == nil {
		if v, ok := this.opts["id"]; ok {
			success = !d.Where("id = ?", v).First(&resource).RecordNotFound()
		}
	}
	return resource, success
}
