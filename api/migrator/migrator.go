package migrator

import (
	"fmt"
	"github.com/confur-me/confur-api/api/models"
	"github.com/confur-me/confur-api/db"
)

func DbDrop() {
	d, err := db.Connection()
	if err != nil {
		fmt.Println("Db connection error", err)
	} else {
		d.DropTable(&models.Conference{})
		d.DropTable(&models.Event{})
		d.DropTable(&models.Video{})
		d.DropTable(&models.Tag{})
		d.DropTable(&models.User{})
	}
}

func DbCreate() {
	d, err := db.Connection()
	if err != nil {
		fmt.Println("Db connection error", err)
	} else {
		d.CreateTable(&models.Conference{})
		d.CreateTable(&models.Event{})
		d.CreateTable(&models.Video{})
		d.CreateTable(&models.Tag{})
		d.CreateTable(&models.User{})
	}
}

func DbMigrate() {
	d, err := db.Connection()
	if err != nil {
		fmt.Println("Db connection error", err)
	} else {
		d.AutoMigrate(
			&models.Conference{},
			&models.Event{},
			&models.Video{},
			&models.Tag{},
			&models.User{},
		)
	}
}
