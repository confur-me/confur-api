package models

import (
	"fmt"
	"github.com/confur-me/confur-api/db"
)

type TagAlias struct {
	ID      uint    `gorm:"primary_key" json:"id"`
	Tag     *Tag    `gorm:"foreignkey:tag_slug" json:"tag,omitempty"`
	TagSlug *string `sql:"not null;index" json:"tag_slug"`
	Title   string  `sql:"not null;index" json:"title"`
}

type tagAliasService struct {
	Service
}

func NewTagAliasService(params map[string]interface{}) *tagAliasService {
	s := new(tagAliasService)
	s.params = params
	return s
}

func (this *tagAliasService) TagAlias() (*TagAlias, error) {
	var (
		resource TagAlias
		err      error
	)
	if conn, ok := db.Connection(); ok {
		if v, ok := this.params["tag_alias_id"]; ok {
			err = conn.Where("id = ?", v).First(&resource).Error
		}
	}
	return &resource, err
}

func (this *tagAliasService) TagAliases() (*[]TagAlias, error) {
	var err error
	collection := make([]TagAlias, 0)
	if conn, ok := db.Connection(); ok {
		query := &conn
		if v, ok := this.params["query"]; ok {
			// FIXME: CHECK injection possibility
			query = query.Where("title ILIKE ?", fmt.Sprintf("%%%v%%", v))
		}
		if v, ok := this.params["tag"]; ok {
			var tag Tag
			conn.Find(&tag, "slug = ?", v)
			if tag.Slug != "" {
				err = query.Model(&tag).Related(&collection, "TagAliases").Error
			}
		} else {
			err = query.Find(&collection).Error
		}
	}
	return &collection, err
}
