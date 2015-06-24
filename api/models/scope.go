package models

import (
	"github.com/confur-me/confur-api/db"
)

type Scope struct {
	Slug             string `gorm:"primary_key" json:"slug"`
	Title            string `json:"title"`
	ConferencesCount uint   `json:"conferences_count"`
	VideosCount      uint   `json:"videos_count"`
	EventsCount      uint   `json:"events_count"`
	SpeakersCount    uint   `json:"speakers_count"`
}

type scopeService struct {
	Service
}

func NewScopeService(params map[string]interface{}) *scopeService {
	s := new(scopeService)
	s.params = params
	return s
}

func (this *scopeService) Scopes() (*[]Scope, error) {
	var (
		collection []Scope = make([]Scope, 0)
		err        error
	)
	if conn, ok := db.Connection(); ok {
		err = conn.Find(&collection).Error
	}
	return &collection, err
}

func (this *scopeService) Scope() (*Scope, error) {
	var (
		resource Scope
		err      error
	)
	if conn, ok := db.Connection(); ok {
		if v, ok := this.params["scope"]; ok {
			err = conn.Where("slug = ?", v).First(&resource).Error
		}
	}
	return &resource, err
}
