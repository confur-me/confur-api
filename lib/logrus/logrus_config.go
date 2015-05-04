package logrus_config

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func init() {
	env := os.Getenv("CONFUR_ENV")
	if env != "production" {
		env = "development"
	}
	if env == "production" {
		// Log as JSON instead of the default ASCII formatter.
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.WarnLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}
	log.SetOutput(os.Stdout)
	// TODO: in development use io.Writer logger
}
