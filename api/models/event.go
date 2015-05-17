package models

import (
	"github.com/confur-me/confur-api/db"
	"time"
)

type Event struct {
	ID             uint     `gorm:"primary_key"`
	ConferenceSlug string   `sql:"index" binding:"required"`
	Title          string   `sql:"type:text" binding:"required"`
	Description    string   `sql:"type:text"`
	Country        string   `sql:"index:idx_country_state_city_address"`
	City           string   `sql:"index:idx_country_state_city_address"`
	State          string   `sql:"index:idx_country_state_city_address"`
	Address        string   `sql:"type:text;index:idx_country_state_city_address"`
	Authors        []Author `gorm:"many2many:events_authors"`
	UpdatedAt      time.Time
	StartedAt      time.Time `sql:"index"`
	DeletedAt      time.Time
}

func EventsByConference(conferenceId string) []Event {
	var collection []Event = make([]Event, 0)
	d, err := db.Connection()
	if err == nil {
		d.Where("conference_id = ?", conferenceId).Find(&collection)
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
