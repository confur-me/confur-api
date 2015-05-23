package db

import (
	"time"

	log "github.com/Sirupsen/logrus"
	_ "github.com/confur-me/confur-api/lib/logrus"

	"github.com/confur-me/confur-api/lib/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Logger struct{}

func (logger Logger) Print(values ...interface{}) {
	if len(values) > 1 {
		currentTime := time.Now()
		duration := values[2]

		log.WithFields(log.Fields{
			"time": currentTime,
			"db":   duration,
		}).Info(values[3:]...)
	}
}

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
	db_config, err := config.Config().Get("db")
	if err != nil {
		log.Error(err)
		return err
	}
	connection_string := "host=" + db_config.UString("host", "localhost") +
		" port=" + db_config.UString("port", "5432") +
		" user=" + db_config.UString("user", "postgres") +
		" dbname=" + db_config.UString("database", "confur") +
		" sslmode=" + db_config.UString("ssl", "disable") +
		" password=" + db_config.UString("password")

	db.connection, err = gorm.Open("postgres", connection_string)
	if err != nil {
		db.connected = false
		log.Error("Unable to connect to database")
		return err
	}
	db.connected = true
	log.Info("Connected to database")
	db.connection.LogMode(true)
	//db.connection.SetLogger(Logger{})
	return nil
}

func Connection() (gorm.DB, bool) {
	var err error
	if db.connected != true {
		err = Connect()
	}
	return db.connection, err == nil
}
