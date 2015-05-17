package models

import (
	"github.com/confur-me/confur-api/db"
	"time"
)

type Video struct {
	ID             uint   `gorm:"primary_key"`
	Title          string `sql:"type:text" binding:"required"`
	Url            string
	Length         int32
	Description    string `sql:"type:text"`
	Service        string `sql:"index:idx_service_service_id" binding:"required"`
	ServiceID      string `sql:"index:idx_service_service_id" binding:"required"`
	ConferenceSlug string `sql:"index" binding:"required"`
	Tags           []Tag  `gorm:"many2many:videos_tags"`
	AuthorID       uint   `sql:"index"`
	LikesCount     int8
	Thumbnail      string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

// TODO: inject likes count

func Videos(limit int, offset int) []Video {
	var collection []Video = make([]Video, 0)
	d, err := db.Connection()
	if err == nil {
		d.Limit(limit).Offset(offset).Find(&collection)
	}
	return collection
}

func VideosByConference(conferenceSlug string) []Video {
	var collection []Video = make([]Video, 0)
	var conference Conference
	d, err := db.Connection()
	if err == nil {
		d.Find(&conference, "slug = ?", conferenceSlug)
		if conference.Slug != "" {
			d.Where("conference_slug = ?", conference.Slug).Find(&collection)
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
	var collection []Video = make([]Video, 0)
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
