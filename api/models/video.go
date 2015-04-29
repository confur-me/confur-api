package models

import (
	"github.com/jinzhu/gorm"
)

type Video struct {
	gorm.Model
	Title  string
	Url    string
	Length int32
}
