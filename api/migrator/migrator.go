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
		)
	}
}
