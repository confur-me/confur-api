package models

import (
	"fmt"
	"github.com/confur-me/confur-api/db"
	"time"
)

type Tag struct {
	ID          uint    `gorm:"primary_key"`
	Slug        string  `sql:"index" binding:"required"`
	Title       string  `binding:"required"`
	Videos      []Video `gorm:"many2many:videos_tags" json:",omitempty"`
	VideosCount int
	DeletedAt   time.Time `json:",omitempty"`
}

type tagService struct {
	Service
}

func NewTagService(params map[string]interface{}) *tagService {
	s := new(tagService)
	s.params = params
	return s
}

func (this *tagService) Tag() (*Tag, error) {
	var (
		resource Tag
		err      error
	)
	if conn, ok := db.Connection(); ok {
		if v, ok := this.params["tag"]; ok {
			err = conn.Where("slug = ?", v).First(&resource).Error
		}
	}
	return &resource, err
}

func (this *tagService) Tags() (*[]Tag, error) {
	var err error
	collection := make([]Tag, 0)
	if conn, ok := db.Connection(); ok {
		query := &conn
		limit := 20 // Defaults to 20 items per page
		page := 0
		if v, ok := this.params["query"]; ok {
			// FIXME: CHECK injection possibility
			query = query.Where("title ILIKE ?", fmt.Sprintf("%%%v%%", v))
		}
		if v, ok := this.params["limit"]; ok {
			if v.(int) <= 50 {
				limit = v.(int)
			}
			if v, ok := this.params["page"]; ok {
				page = v.(int)
			}
			query = query.Scopes(Paginate(limit, page))
		}
		err = query.Limit(limit).Find(&collection).Error
	}
	return &collection, err
}
