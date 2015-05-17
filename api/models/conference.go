package models

import (
	"github.com/confur-me/confur-api/db"
	"time"
)

type Conference struct {
	Slug        string `sql:"type:varchar(64)" gorm:"primary_key" binding:"required"`
	Title       string `sql:"type:text" binding:"required"`
	Url         string
	Type        string `sql:"index" binding:"required"`
	Description string `sql:"type:text"`
	Events      []Event
	Videos      []Video
	VideosCount int
	Thumbnail   string
	IsActive    bool `sql:"index"`
	UpdatedAt   time.Time
}

func Conferences() []Conference {
	var collection []Conference = make([]Conference, 0)
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
