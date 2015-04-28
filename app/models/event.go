package models

import (
	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	ConferenceID uint `sql:"index"`
	Address      string
}
