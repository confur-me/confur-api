package fixtures

import (
	"github.com/confur-me/confur-api/api/models"
	"github.com/confur-me/confur-api/db"
)

func Seed() {
	if d, ok := db.Connection(); ok {
		tags := []models.Tag{{
			Slug:  "go",
			Title: "Go",
		}, {
			Slug:  "ruby",
			Title: "Ruby",
		}}

		for _, tag := range tags {
			d.Create(&tag)
		}

		conferences := []models.Conference{{
			Slug:        "moscowjs",
			Title:       "MoscowJS",
			Url:         "http://moscowjs.ru",
			Type:        "meetup",
			Events:      []models.Event{{Address: "Russian Federation, Moscow, Red square street."}},
			Description: "Cum illum voluptas ducimus",
			Videos: []models.Video{{
				Title:   "Videooo",
				Url:     "http://www.youtube.com/watch?v=oHg5SJYRHA0",
				Length:  180,
				Service: "youtube",
				Tags:    tags,
			}, {
				Title:   "Videooo 2",
				Url:     "http://www.youtube.com/watch?v=oHg5SJYRHA0",
				Length:  120,
				Service: "youtube",
			}},
		},
			{
				Slug:   "railsclub",
				Title:  "Rails Club",
				Url:    "http://gay.org",
				Type:   "conference",
				Events: []models.Event{{Address: "Russian Federation, Siberia."}},
			}}

		for _, conference := range conferences {
			d.Create(&conference)
		}
	}

}
