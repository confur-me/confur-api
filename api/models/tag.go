package models

import (
	"fmt"
	"github.com/confur-me/confur-api/db"
	"time"
)

type Tag struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Slug        string    `sql:"index" binding:"required" json:"slug"`
	Title       string    `binding:"required" json:"title"`
	Videos      []Video   `gorm:"many2many:videos_tags" json:"videos,omitempty"`
	VideosCount int       `json:"videos_count"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
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
		page := 1
		if v, ok := this.params["query"]; ok {
			// FIXME: CHECK injection possibility
			query = query.Where("slug ILIKE ?", fmt.Sprintf("%%%v%%", v))
		}
		if v, ok := this.params["limit"]; ok {
			if v.(int) <= 50 {
				limit = v.(int)
			}
		}
		if v, ok := this.params["page"]; ok {
			page = v.(int)
		}
		err = query.Scopes(Paginate(limit, page)).Limit(limit).Find(&collection).Error
	}
	return &collection, err
}
