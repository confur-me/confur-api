package models

import (
	"fmt"
	"github.com/confur-me/confur-api/db"
)

type Tag struct {
	ID          uint    `gorm:"primary_key"`
	Slug        string  `sql:"index" binding:"required"`
	Title       string  `binding:"required"`
	Videos      []Video `gorm:"many2many:videos_tags"`
	VideosCount int
}

type TagService struct {
	service
}

func NewTagService(opts map[string]interface{}) *TagService {
	s := new(TagService)
	s.opts = opts
	return s
}

func (this *TagService) FindTag() (Tag, bool) {
	var (
		resource Tag
		success  bool
	)
	if d, err := db.Connection(); err == nil {
		if v, ok := this.opts["tag"]; ok {
			success = !d.Where("slug = ?", v).First(&resource).RecordNotFound()
		}
	}
	return resource, success
}

func (this *TagService) FindTags() []Tag {
	collection := make([]Tag, 0)
	if d, err := db.Connection(); err == nil {
		query := &d
		limit := 20 // Defaults to 20 items per page
		if v, ok := this.opts["query"]; ok {
			// FIXME: CHECK injection possibility
			query = query.Where("title ILIKE ?", fmt.Sprintf("%%%v%%", v))
		}
		if v, ok := this.opts["limit"]; ok {
			if v.(int) <= 50 {
				limit = v.(int)
			}
		}
		if v, ok := this.opts["page"]; ok {
			offset := (v.(int) - 1) * limit
			query = query.Offset(offset)
		}
		query = query.Limit(limit).Find(&collection)
	}
	return collection
}
