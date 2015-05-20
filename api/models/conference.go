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
	Videos      []Video `json:",omitempty"`
	VideosCount int
	Thumbnail   string
	IsActive    bool      `sql:"index"`
	UpdatedAt   time.Time `json:",omitempty"`
}

type ConferenceService struct {
	Service
}

func NewConferenceService(opts map[string]interface{}) *ConferenceService {
	s := new(ConferenceService)
	s.opts = opts
	return s
}

func (this *ConferenceService) FindConferences() []Conference {
	var collection []Conference = make([]Conference, 0)
	if d, ok := db.Connection(); ok {
		query := &d
		limit := 20 // Defaults to 20 items per page
		if v, ok := this.opts["query"]; ok {
			// FIXME: CHECK injection possibility
			query = query.Where("title ILIKE ?", "%"+fmt.Sprintf("%v", v)+"%")
		}
		if v, ok := this.opts["limit"]; ok {
			if v.(int) <= 50 {
				limit = v.(int)
			}
		}
		if v, ok := this.opts["page"]; ok {
			offset := (v.(int) - 1) * limit
			query = query.Offset(offset)
		}
		query = query.Limit(limit).Find(&collection)
	}
	return collection
}

func (this *ConferenceService) FindConference() (Conference, bool) {
	var (
		resource Conference
		success  bool
	)
	if d, ok := db.Connection(); ok {
		if v, ok := this.opts["conference"]; ok {
			success = !d.Where("slug = ?", v).First(&resource).RecordNotFound()
		}
	}
	return resource, success
}
