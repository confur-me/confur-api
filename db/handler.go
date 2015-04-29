package db

import (
	//"github.com/confur-me/confur-api/lib/config"
	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type DB struct {
	connection gorm.DB
	connected  bool
}

func New() *DB {
	db := new(DB)
	return db
}

var db *DB

func init() {
	db = New()
}

func Connect() error {
	//db_config := config.GetStringMapString("db")
	var err error
	db.connection, err = gorm.Open("postgres", "user=postgres dbname=confur sslmode=disable")
	if err != nil {
		db.connected = false
		return err
	}
	db.connected = true
	db.connection.LogMode(true)
	return nil
}

func Connection() (gorm.DB, error) {
	var err error
	if db.connected != true {
		err = Connect()
	}
	return db.connection, err
}
