package models

import (
	"github.com/confur-me/confur-api/db"
	"github.com/jinzhu/gorm"
	"time"
)

type Event struct {
	gorm.Model
	ConferenceID uint   `sql:"index"`
	Address      string `sql:"type:text"`
	Description  string `sql:"type:text"`
	StartedAt    *time.Time
}

func ConferenceEventsCollection(conference_id string) []Event {
	var collection []Event
	d, err := db.Connection()
	if err == nil {
		d.Where("conference_id = ?", conference_id).Find(&collection)
	}
	return collection
}

func EventById(id string) Event {
	var resource Event
	d, err := db.Connection()
	if err == nil {
		d.Where("id = ?", id).First(&resource)
	}
	return resource
}
