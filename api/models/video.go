package models

import (
	"fmt"
	db "github.com/confur-me/confur-api/db"
	"time"
)

type Video struct {
	ID             uint      `gorm:"primary_key" json:"id"`
	Title          string    `sql:"type:text" binding:"required" json:"title"`
	Url            string    `json:"url"`
	Length         int32     `json:"length"`
	Description    string    `sql:"type:text" json:"description"`
	Service        string    `sql:"index:idx_service_service_id" binding:"required" json:"service"`
	ServiceID      string    `sql:"index:idx_service_service_id" binding:"required" json:"service_id"`
	ConferenceSlug string    `sql:"index" binding:"required" json:"conference_slug"`
	Tags           []Tag     `gorm:"many2many:videos_tags" json:"tags,omitempty"`
	Speakers       []Speaker `gorm:"many2many:videos_speakers" json:"speakers,omitempty"`
	EventID        uint      `sql:"index" json:"event_id"`
	LikesCount     int8      `json:"likes_count"`
	Thumbnail      string    `sql:"type:text" json:"thumbnail"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at,omitempty"`
}

type videoService struct {
	Service
}

func NewVideoService(params map[string]interface{}) *videoService {
	s := new(videoService)
	s.params = params
	return s
}

// TODO: inject likes count

func (this *videoService) Videos() (*[]Video, error) {
	var err error
	collection := make([]Video, 0)
	limit := 20
	page := 1
	if conn, ok := db.Connection(); ok {
		query := &conn
		//if v, ok := this.params["tag"]; ok {
		//var tag Tag
		//d.Find(&tag, "slug = ?", v)
		//if tag.ID > 0 {
		////d.Model(&tag).Related(&collection, "Videos")
		////query = query.Where("conference_slug = ?", v)
		//}
		//}
		if v, ok := this.params["conference"]; ok {
			query = query.Where("conference_slug = ?", v)
		}
		if v, ok := this.params["event"]; ok {
			query = query.Where("event_id = ?", v)
		}
		if v, ok := this.params["query"]; ok {
			// FIXME: CHECK injection possibility
			query = query.
				Where("title ILIKE ?", fmt.Sprintf("%%%v%%", v))
		}
		if v, ok := this.params["limit"]; ok {
			if v.(int) <= 50 {
				limit = v.(int)
			}
		}
		if v, ok := this.params["page"]; ok {
			page = v.(int)
		}
		err = query.Scopes(Paginate(limit, page)).Find(&collection).Error
	}
	return &collection, err
}

func (this *videoService) Video() (*Video, error) {
	var (
		resource Video
		err      error
		tags     []Tag
		speakers []Speaker
	)
	if conn, ok := db.Connection(); ok {
		if v, ok := this.params["video"]; ok {
			err = conn.Where("id = ?", v).First(&resource).Error
			if err == nil {
				conn.Model(&resource).Related(&tags, "Tags").Related(&speakers, "Speakers")
				resource.Tags = tags
				resource.Speakers = speakers
			}
		}
	}
	return &resource, err
}
