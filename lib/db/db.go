package db

import (
	"github.com/confur-me/confur-api/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Db struct {
	// TODO: pass config here
}

func (this *Db) Connect() error {
	// TODO: return connection handler
	return nil
}

func (this *Db) Setup() error {
	return nil
}

func (this *Db) Migrate() error {
	db, err := gorm.Open("postgres", "user=gorm dbname=gorm sslmode=disable")
	if err == nil {
		db.AutoMigrate(
			&models.Conference{},
			&models.Event{},
			&models.Video{},
		)
	}
	return err
}
