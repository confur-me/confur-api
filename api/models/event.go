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
	Authors        []Author `gorm:"many2many:events_authors" json:",omitempty"`
	UpdatedAt      time.Time
	StartedAt      time.Time `sql:"index"`
	DeletedAt      time.Time
}

type EventService struct {
	service
}

func NewEventService(opts map[string]interface{}) *EventService {
	s := new(EventService)
	s.opts = opts
	return s
}

func (this *EventService) FindEvent() (Event, bool) {
	var (
		resource Event
		success  bool
	)
	if d, ok := db.Connection(); ok {
		if v, ok := this.opts["event"]; ok {
			success = !d.Where("id = ?", v).First(&resource).RecordNotFound()
		}
	}
	return resource, success
}

func (this *EventService) FindEvents() []Event {
	collection := make([]Event, 0)
	if d, ok := db.Connection(); ok {
		if v, ok := this.opts["conference"]; ok {
			d.Where("conference_slug = ?", v).Find(&collection)
		}
	}
	return collection
}
