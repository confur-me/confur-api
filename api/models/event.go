package models

import (
	"github.com/confur-me/confur-api/db"
	"time"
)

type Event struct {
	ID           uint `gorm:"primary_key"`
	ConferenceID uint `sql:"index"`
	Country      string
	City         string
	State        string
	Address      string `sql:"type:text"`
	Description  string `sql:"type:text"`
	StartedAt    time.Time
}

func EventsByConference(conference_id string) []Event {
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
