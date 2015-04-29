package models

import (
	"github.com/confur-me/confur-api/db"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Slug   string `sql:"index"`
	Title  string
	Videos []Video `gorm:"many2many:video_tags"`
}

func TagBySlug(slug string) Tag {
	var resource Tag
	d, err := db.Connection()
	if err == nil {
		d.Where("slug = ?", slug).First(&resource)
	}
	return resource
}

func TagsCollection() []Tag {
	var collection []Tag
	d, err := db.Connection()
	if err == nil {
		d.Find(&collection)
	}
	return collection
}
