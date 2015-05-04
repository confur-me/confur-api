package models

import (
	"github.com/confur-me/confur-api/db"
	"time"
)

type Conference struct {
	ID          uint   `gorm:"primary_key"`
	Slug        string `sql:"index"`
	Title       string
	Url         string
	Type        string `sql:"index"`
	Description string `sql:"type:text"`
	Events      []Event
	Videos      []Video
	VideosCount int
	UpdatedAt   time.Time
}

func Conferences() []Conference {
	var collection []Conference
	d, err := db.Connection()
	if err == nil {
		d.Find(&collection)
	}
	return collection
}

func ConferenceBySlug(slug string) Conference {
	var resource Conference
	d, err := db.Connection()
	if err == nil {
		d.Where("slug = ?", slug).First(&resource)
	}
	return resource
}
