package models

import (
	"fmt"
	db "github.com/confur-me/confur-api/db"
	"time"
)

type Video struct {
	ID             uint        `gorm:"primary_key" json:"id"`
	Title          *string     `sql:"not null,type:text" binding:"required" json:"title"`
	Url            string      `json:"url"`
	Length         *uint       `json:"length,omitempty"`
	Description    string      `sql:"type:text" json:"description"`
	ConferenceSlug *string     `sql:"not null,index" binding:"required" json:"conference_slug"`
	Conference     *Conference `json:"conference,omitempty" gorm:"foreignkey:conference_slug"`
	Scope          *string     `sql:"not null,index" json:"scope"`
	TagResources   *[]Tag      `gorm:"many2many:videos_tags;associationforeignkey:tag_slug" json:"-"`
	Tags           []string    `sql:"-" json:"tags"`
	Speakers       *[]Speaker  `gorm:"many2many:videos_speakers" json:"speakers,omitempty"`
	EventID        *uint       `sql:"index" json:"event_id"`
	Event          *Event      `json:"event,omitempty" gorm:"foreignkey:event_id"`
	LikesCount     uint        `json:"likes_count"`
	DislikesCount  uint        `json:"likes_count"`
	Rating         float64     `json:"rating"`
	Thumbnail      string      `sql:"type:text" json:"thumbnail"`
	Language       *string     `sql:"type:varchar(2)" json:"language"`
	DeletedAt      *time.Time  `json:"deleted_at,omitempty"`
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

func (this *videoService) Videos() (*[]Video, int, int, int, error) {
	var (
		err        error
		count      int
		offset     int
		collection []Video = make([]Video, 0)
		limit      int     = 20
	)
	if conn, ok := db.Connection(); ok {
		query := &conn
		if v, ok := this.params["conference"]; ok {
			query = query.Where("conference_slug = ?", v)
		}
		if v, ok := this.params["event"]; ok {
			query = query.Where("event_id = ?", v)
		}
		if v, ok := this.params["query"]; ok {
			// FIXME: CHECK sql injection possibility
			query = query.
				Where("title ILIKE ?", fmt.Sprintf("%%%v%%", v))
		}
		if _, ok := this.params["shuffle"]; ok {
			query = query.Where("random() < 0.01")
		}
		if v, ok := this.params["limit"]; ok {
			if v.(int) > 0 && v.(int) <= 50 {
				limit = v.(int)
			}
		}
		if v, ok := this.params["offset"]; ok {
			offset = v.(int)
		}
		if v, ok := this.params["tag"]; ok {
			query = query.Joins("INNER JOIN videos_tags ON videos_tags.video_id = videos.id").Where("videos_tags.tag_slug = ?", v)
		}

		err = query.
			Preload("Conference").
			Preload("Event").
			Scopes(GetRange(limit, offset)).
			Find(&collection).
			Scopes(Unpaginate).
			Count(&count).
			Error

		if err == nil {
			tags := make(map[uint][]string)
			videoIds := make([]uint, 0)
			for _, video := range collection {
				videoIds = append(videoIds, video.ID)
			}
			rows, err1 := conn.Table("videos_tags").Select("video_id, tag_slug").Where("video_id IN (?)", videoIds).Rows()
			if err1 == nil {
				var (
					videoId uint
					tag     string
				)
				for rows.Next() {
					rows.Scan(&videoId, &tag)
					tags[videoId] = append(tags[videoId], tag)
				}
				for i, video := range collection {
					video.Tags = tags[video.ID]
					collection[i] = video
				}
			}
		}
	}
	return &collection, count, limit, offset, err
}

func (this *videoService) Video() (*Video, error) {
	var (
		resource Video
		err      error
	)
	if conn, ok := db.Connection(); ok {
		if v, ok := this.params["video"]; ok {
			err = conn.Where("id = ?", v).Preload("Conference").Preload("Event").First(&resource).Error
			if err == nil {
				var tagResources []Tag
				speakers := make([]Speaker, 0)
				tags := make([]string, 0)

				conn.Model(&resource).Related(&tagResources, "TagResources").Related(&speakers, "Speakers")
				for _, v := range tagResources {
					tags = append(tags, v.Slug)
				}
				resource.Tags = tags
				resource.Speakers = &speakers
			}
		}
	}
	return &resource, err
}
