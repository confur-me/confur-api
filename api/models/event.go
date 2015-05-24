package models

import (
	"github.com/confur-me/confur-api/db"
	"time"
)

type Event struct {
	ID             uint      `gorm:"primary_key" json:"id"`
	ConferenceSlug string    `sql:"index" binding:"required" json:"conference_slug"`
	Title          string    `sql:"type:text" binding:"required" json:"title"`
	Description    string    `sql:"type:text" json:"description"`
	Country        string    `sql:"index:idx_country_state_city_address" json:"country"`
	City           string    `sql:"index:idx_country_state_city_address" json:"city"`
	State          string    `sql:"index:idx_country_state_city_address" json:"state"`
	Address        string    `sql:"type:text;index:idx_country_state_city_address" json:"address"`
	Authors        []Author  `gorm:"many2many:events_authors" json:"authors,omitempty"`
	VideosCount    uint      `sql:"not null;default:0" json:"videos_count"`
	UpdatedAt      time.Time `json:"updated_at"`
	StartedAt      time.Time `sql:"index" json:"started_at"`
	DeletedAt      time.Time `json:"deleted_at,omitempty"`
}

type eventService struct {
	Service
}

func NewEventService(params map[string]interface{}) *eventService {
	s := new(eventService)
	s.params = params
	return s
}

func (this *eventService) Event() (*Event, error) {
	var (
		resource Event
		err      error
	)
	if conn, ok := db.Connection(); ok {
		if v, ok := this.params["event"]; ok {
			err = conn.Where("id = ?", v).First(&resource).Error
		}
	}
	return &resource, err
}

func (this *eventService) Events() (*[]Event, error) {
	var err error
	collection := make([]Event, 0)
	if conn, ok := db.Connection(); ok {
		if v, ok := this.params["conference"]; ok {
			err = conn.Where("conference_slug = ?", v).Find(&collection).Error
		}
	}
	return &collection, err
}
