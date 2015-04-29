package db

import (
	"fmt"
	"github.com/confur-me/confur-api/lib/config"
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
	db_config := config.GetStringMapString("db")
	connection_string := "host=" + db_config["host"] +
		" port=" + db_config["port"] +
		" user=" + db_config["user"] +
		" dbname=" + db_config["database"] +
		" sslmode=" + db_config["ssl"] +
		" password=" + db_config["password"]

	var err error
	db.connection, err = gorm.Open("postgres", connection_string)
	if err != nil {
		db.connected = false
		fmt.Println("Unable to connect to database")
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
