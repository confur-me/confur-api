package models

import (
	"fmt"
	"github.com/confur-me/confur-api/db"
	"time"
)

type Event struct {
	ID             uint        `gorm:"primary_key" json:"id"`
	ConferenceSlug *string     `sql:"not null;index" binding:"required" json:"conference_slug"`
	Conference     *Conference `json:"conference,omitempty" gorm:"foreignkey:conference_slug"`
	Scope          *string     `sql:"not null;index" json:"scope" binding:"required"`
	Title          string      `sql:"type:text" binding:"required" json:"title"`
	Description    string      `sql:"type:text" json:"description"`
	Country        string      `sql:"index:idx_country_state_city_address" json:"country"`
	City           string      `sql:"index:idx_country_state_city_address" json:"city"`
	State          string      `sql:"index:idx_country_state_city_address" json:"state"`
	Address        string      `sql:"type:text;index:idx_country_state_city_address" json:"address"`
	Speakers       *[]Speaker  `gorm:"many2many:events_speakers" json:"speakers,omitempty"`
	Videos         *[]Video    `json:"videos,omitempty"`
	VideosCount    uint        `sql:"not null;default:0" json:"videos_count"`
	Thumbnail      string      `json:"thumbnail"`
	IsActive       *bool       `sql:"not null;index" binding:"required" json:"-"`
	StartsAt       *time.Time  `sql:"index" json:"starts_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	DeletedAt      *time.Time  `json:"deleted_at,omitempty"`
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
			err = conn.Scopes(Active).Where("id = ?", v).Preload("Conference").First(&resource).Error
		}
	}
	return &resource, err
}

func (this *eventService) Events() (*[]Event, error) {
	var err error
	collection := make([]Event, 0)
	if conn, ok := db.Connection(); ok {
		query := &conn
		limit := 20 // Defaults to 20 items per page
		page := 1
		if v, ok := this.params["conference"]; ok {
			query = query.Where("conference_slug = ?", v)
		}
		if v, ok := this.params["query"]; ok {
			// FIXME: CHECK injection possibility
			query = query.Where("title ILIKE ?", fmt.Sprintf("%%%v%%", v))
		}
		if _, ok := this.params["shuffle"]; ok {
			query = query.Where("random() < 0.1")
		}
		if v, ok := this.params["limit"]; ok {
			if v.(int) <= 50 {
				limit = v.(int)
			}
		}
		if v, ok := this.params["page"]; ok {
			page = v.(int)
		}
		err = query.Scopes(Active, Paginate(limit, page)).Preload("Conference").Find(&collection).Error
	}
	return &collection, err
}
