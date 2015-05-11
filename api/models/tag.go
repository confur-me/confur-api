package models

import (
	"github.com/confur-me/confur-api/db"
)

type Tag struct {
	ID          uint   `gorm:"primary_key"`
	Slug        string `sql:"index"`
	Title       string
	Videos      []Video `gorm:"many2many:videos_tags"`
	VideosCount int
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
	var collection []Tag = make([]Tag, 0)
	d, err := db.Connection()
	if err == nil {
		d.Find(&collection)
	}
	return collection
}
