package migrator

import (
	"github.com/confur-me/confur-api/api/models"
	"github.com/confur-me/confur-api/db"
	_ "github.com/confur-me/confur-api/lib/logrus"
)

func DbDrop() {
	if d, ok := db.Connection(); ok {
		d.DropTable(&models.Conference{})
		d.DropTable(&models.Event{})
		d.DropTable(&models.Video{})
		d.DropTable(&models.Tag{})
		d.DropTable(&models.User{})
		d.DropTable(&models.Speaker{})
	}
}

func DbCreate() {
	if d, ok := db.Connection(); ok {
		d.CreateTable(&models.Conference{})
		d.CreateTable(&models.Event{})
		d.CreateTable(&models.Video{})
		d.CreateTable(&models.Tag{})
		d.CreateTable(&models.User{})
		d.CreateTable(&models.Speaker{})
	}
}

func DbMigrate() {
	if d, ok := db.Connection(); ok {
		d.AutoMigrate(
			&models.Conference{},
			&models.Event{},
			&models.Video{},
			&models.Tag{},
			&models.User{},
			&models.Speaker{},
		)
	}
}
