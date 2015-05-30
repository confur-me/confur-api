package models

import (
	"github.com/confur-me/confur-api/db"
	"time"
)

type Speaker struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	MiddleName string    `json:"middle_name"`
	CreatedAt  time.Time `json:"created_at"`
	Photo      string    `json:"photo"`
	Events     []Event   `gorm:"many2many:events_speakers" json:"events,omitempty"`
	Videos     []Video   `gorm:"many2many:videos_speakers" json:"videos,omitempty"`
}

type SpeakerService struct {
	Service
}

func (this *SpeakerService) Speaker() (*Speaker, error) {
	var (
		resource Speaker
		err      error
	)
	if conn, ok := db.Connection(); ok {
		if v, ok := this.params["id"]; ok {
			err = conn.Where("id = ?", v).First(&resource).Error
		}
	}
	return &resource, err
}
