package models

import (
	"fmt"
	"github.com/confur-me/confur-api/db"
	"time"
)

type Conference struct {
	Slug        string     `sql:"type:varchar(64)" gorm:"primary_key" binding:"required" json:"slug"`
	Title       string     `sql:"type:text" binding:"required" json:"title"`
	Url         string     `json:"url"`
	Type        *string    `sql:"not null;index" binding:"required" json:"type"`
	Scope       *string    `sql:"not null;index" json:"scope"`
	Description string     `sql:"type:text" json:"description"`
	Events      []Event    `json:"events,omitempty" gorm:"foreignkey:conference_slug"`
	EventsCount uint       `sql:"not null;default:0" json:"events_count"`
	Videos      []Video    `json:"videos,omitempty" gorm:"foreignkey:conference_slug"`
	VideosCount uint       `sql:"not null;default:0" json:"videos_count"`
	Thumbnail   string     `json:"thumbnail"`
	IsActive    *bool      `sql:"not null;index" binding:"required" json:"-"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type conferenceService struct {
	Service
	Resource *Conference
}

func NewConferenceService(params map[string]interface{}) *conferenceService {
	s := new(conferenceService)
	s.params = params
	return s
}

func (this *conferenceService) Conference() (*Conference, error) {
	var (
		resource Conference
		err      error
	)
	if conn, ok := db.Connection(); ok {
		if v, ok := this.params["conference"]; ok {
			err = conn.Scopes(Active).Where("slug = ?", v).First(&resource).Error
		}
	}
	return &resource, err
}

func (this *conferenceService) Conferences() (*[]Conference, error) {
	var (
		collection []Conference = make([]Conference, 0)
		err        error
	)
	if conn, ok := db.Connection(); ok {
		query := &conn
		limit := 20 // Defaults to 20 items per page
		page := 1
		if v, ok := this.params["query"]; ok {
			// FIXME: CHECK injection possibility
			query = query.Where("title ILIKE ?", fmt.Sprintf("%%%v%%", v))
		}
		if v, ok := this.params["limit"]; ok {
			if v.(int) <= 50 {
				limit = v.(int)
			}
		}
		if v, ok := this.params["page"]; ok {
			page = v.(int)
		}
		if _, ok := this.params["shuffle"]; ok {
			query = query.Where("random() < 0.1")
		}
		err = query.Scopes(Active, Paginate(limit, page)).Find(&collection).Error
	}
	return &collection, err
}
