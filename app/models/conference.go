package models

import (
	"github.com/jinzhu/gorm"
)

type Conference struct {
	gorm.Model
	Title  string
	Url    string
	Type   string `sql:"index"`
	Events []Event
}
