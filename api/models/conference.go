package models

import (
	"fmt"
	"github.com/confur-me/confur-api/db"
	"time"
)

type Conference struct {
	Slug        string `sql:"type:varchar(64)" gorm:"primary_key" binding:"required"`
	Title       string `sql:"type:text" binding:"required"`
	Url         string
	Type        string  `sql:"index" binding:"required"`
	Description string  `sql:"type:text"`
	Events      []Event `json:",omitempty"`
	EventsCount uint    `sql:"not null;default:0"`
	Videos      []Video `json:",omitempty"`
	VideosCount uint    `sql:"not null;default:0"`
	Thumbnail   string
	IsActive    bool `sql:"index"`
	UpdatedAt   time.Time
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
			err = conn.Where("slug = ?", v).First(&resource).Error
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
		page := 0
		if v, ok := this.params["query"]; ok {
			// FIXME: CHECK injection possibility
			query = query.Where("title ILIKE ?", fmt.Sprintf("%%%v%%", v))
		}
		if v, ok := this.params["limit"]; ok {
			if v.(int) <= 50 {
				limit = v.(int)
			}
			if v, ok := this.params["page"]; ok {
				page = v.(int)
			}
			query = query.Scopes(Paginate(limit, page))
		}
		err = query.Find(&collection).Error
	}
	return &collection, err
}
