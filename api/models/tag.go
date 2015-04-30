package models

import (
	"github.com/confur-me/confur-api/db"
)

type Tag struct {
	ID     uint   `sql:"primary_key"`
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

func Tags() []Tag {
	var collection []Tag
	d, err := db.Connection()
	if err == nil {
		d.Find(&collection)
	}
	return collection
}
