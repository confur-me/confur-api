package fixtures

import (
	"fmt"
	"github.com/confur-me/confur-api/api/models"
	"github.com/confur-me/confur-api/db"
)

func Seed() {
	d, err := db.Connection()
	if err != nil {
		fmt.Println("Db connection error", err)
	} else {
		conference := models.Conference{
			Title: "MoscowJS",
			Url:   "http://moscowjs.ru",
			Type:  "meetup",
		}

		d.Create(&conference)
	}

}
