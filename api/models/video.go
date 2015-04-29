package models

import (
	"github.com/confur-me/confur-api/db"
	"github.com/jinzhu/gorm"
)

type Video struct {
	gorm.Model
	Title        string
	Url          string
	Length       int32
	Service      string `sql:"index"`
	ConferenceID uint   `sql:"index"`
}

func ConferenceVideosCollection(conference_id string) []Video {
	var collection []Video
	d, err := db.Connection()
	if err == nil {
		d.Where("conference_id = ?", conference_id).Find(&collection)
	}
	return collection
}

func VideoById(id string) Video {
	var resource Video
	d, err := db.Connection()
	if err == nil {
		d.Where("id = ?", id).First(&resource)
	}
	return resource
}
