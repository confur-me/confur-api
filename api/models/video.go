package models

import (
	"fmt"
	"github.com/confur-me/confur-api/db"
	"time"
)

type Video struct {
	ID             uint   `gorm:"primary_key"`
	Title          string `sql:"type:text" binding:"required"`
	Url            string
	Length         int32
	Description    string `sql:"type:text"`
	Service        string `sql:"index:idx_service_service_id" binding:"required"`
	ServiceID      string `sql:"index:idx_service_service_id" binding:"required"`
	ConferenceSlug string `sql:"index" binding:"required"`
	Tags           []Tag  `gorm:"many2many:videos_tags"`
	AuthorID       uint   `sql:"index"`
	LikesCount     int8
	Thumbnail      string `sql:"type:text"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

type VideoService struct {
	service
}

func NewVideoService(opts map[string]interface{}) *VideoService {
	s := new(VideoService)
	s.opts = opts
	return s
}

// TODO: inject likes count

func (this *VideoService) FindVideos() []Video {
	collection := make([]Video, 0)
	limit := 20
	if d, err := db.Connection(); err == nil {
		query := &d
		//if v, ok := this.opts["tag"]; ok {
		//var tag Tag
		//d.Find(&tag, "slug = ?", v)
		//if tag.ID > 0 {
		////d.Model(&tag).Related(&collection, "Videos")
		////query = query.Where("conference_slug = ?", v)
		//}
		//}
		if v, ok := this.opts["conference"]; ok {
			query = query.Where("conference_slug = ?", v)
		}
		if v, ok := this.opts["query"]; ok {
			// FIXME: CHECK injection possibility
			query = query.
				Where("title ILIKE ? OR description ILIKE ?", fmt.Sprintf("%%%v%%", v), fmt.Sprintf("%%%v%%", v))
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

func (this *VideoService) FindVideo() (Video, bool) {
	var (
		resource Video
		success  bool
	)
	if d, err := db.Connection(); err == nil {
		if v, ok := this.opts["id"]; ok {
			success = !d.Where("id = ?", v).First(&resource).RecordNotFound()
		}
	}
	return resource, success
}
