package models

import (
	"github.com/confur-me/confur-api/db"
	"github.com/jinzhu/gorm"
)

type Conference struct {
	gorm.Model
	Title  string
	Url    string
	Type   string `sql:"index"`
	Events []Event
}

var resource *Conference

func ConferencesCollection() []Conference {
	var collection []Conference
	d, err := db.Connection()
	if err == nil {
		d.Find(&collection)
	}
	return collection
}
