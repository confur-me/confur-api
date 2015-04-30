package models

import (
	"github.com/confur-me/confur-api/db"
)

type Conference struct {
	ID          uint `gorm:"primary_key"`
	Title       string
	Url         string
	Type        string `sql:"index"`
	Description string `sql:"type:text"`
	Events      []Event
	Videos      []Video
}

func Conferences() []Conference {
	var collection []Conference
	d, err := db.Connection()
	if err == nil {
		d.Find(&collection)
	}
	return collection
}

func ConferenceById(id string) Conference {
	var resource Conference
	d, err := db.Connection()
	if err == nil {
		d.Where("id = ?", id).First(&resource)
	}
	return resource
}
