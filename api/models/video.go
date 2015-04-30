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
	Tags         []Tag  `gorm:"many2many:video_tags"`
}

func VideosByConference(conferenceId string) []Video {
	var collection []Video
	d, err := db.Connection()
	if err == nil {
		d.Where("conference_id = ?", conferenceId).Find(&collection)
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

func VideosByTag(tagSlug string) []Video {
	var collection []Video
	var tag Tag
	d, err := db.Connection()
	if err == nil {
		d.Find(&tag, "slug = ?", tagSlug)
		if tag.ID > 0 {
			d.Model(&tag).Related(&collection, "Videos")
		}
	}
	return collection
}
