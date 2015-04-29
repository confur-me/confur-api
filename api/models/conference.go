package models

import (
	//"github.com/confur-me/confur-api/core"
	"github.com/jinzhu/gorm"
)

type Conference struct {
	gorm.Model
	Title  string
	Url    string
	Type   string `sql:"index"`
	Events []Event
}

var collection *[]Conference
var resource *Conference

func ConferencesCollection() *[]Conference {
	return collection
}
