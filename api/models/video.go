package models

import (
	"github.com/confur-me/confur-api/db"
	"time"
)

type Video struct {
	ID           uint   `gorm:"primary_key"`
	Title        string `sql:"type:text"`
	Url          string
	Length       int32
	Description  string `sql:"type:text"`
	Service      string `sql:"index:idx_service_service_id"`
	ServiceID    string `sql:"index:idx_service_service_id"`
	ConferenceID uint   `sql:"index"`
	Tags         []Tag  `gorm:"many2many:videos_tags"`
	LikesCount   int8
	CreatedAt    time.Time
	UpdatedAt    time.Time
	AuthorID     uint `sql:"index"`
}

// TODO: inject likes count

func VideosByConference(conferenceSlug string) []Video {
	var collection []Video
	var conference Conference
	d, err := db.Connection()
	if err == nil {
		d.Find(&conference, "slug = ?", conferenceSlug)
		if conference.ID > 0 {
			d.Where("conference_id = ?", conference.ID).Find(&collection)
		}
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
