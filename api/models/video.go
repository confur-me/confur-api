package models

import (
	"fmt"
	db "github.com/confur-me/confur-api/db"
	"time"
)

type Video struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	Title          *string    `sql:"not null,type:text" binding:"required" json:"title"`
	Url            string     `json:"url"`
	Length         uint       `json:"length"`
	Description    string     `sql:"type:text" json:"description"`
	ConferenceSlug *string    `sql:"not null,index" binding:"required" json:"conference_slug"`
	Scope          *string    `sql:"not null,index" json:"scope"`
	TagResources   []Tag      `gorm:"many2many:videos_tags" json:"-"`
	Tags           []string   `sql:"-" json:"tags,omitempty"`
	Speakers       []Speaker  `gorm:"many2many:videos_speakers" json:"speakers,omitempty"`
	EventID        *uint      `sql:"index" json:"event_id"`
	LikesCount     uint       `json:"likes_count"`
	DislikesCount  uint       `json:"likes_count"`
	Rating         float64    `json:"rating"`
	Thumbnail      string     `sql:"type:text" json:"thumbnail"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
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
		resource     Video
		err          error
		tagResources []Tag
		speakers     []Speaker
		tags         []string
	)
	if conn, ok := db.Connection(); ok {
		if v, ok := this.params["video"]; ok {
			err = conn.Where("id = ?", v).First(&resource).Error
			if err == nil {
				conn.Model(&resource).Related(&tags, "TagResources").Related(&speakers, "Speakers")
				for _, v := range tagResources {
					tags = append(tags, v.Slug)
					resource.Tags = tags
				}
				resource.Speakers = speakers
			}
		}
	}
	return &resource, err
}
