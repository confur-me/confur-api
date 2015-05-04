package config

import (
	log "github.com/Sirupsen/logrus"
	_ "github.com/confur-me/confur-api/lib/logrus"
	cfg "github.com/olebedev/config"
)

var c *cfg.Config

func init() {
	c = new(cfg.Config)
}

func Read(path string) error {
	log.Info("Reading configuration from ", path)
	var err error
	c, err = cfg.ParseYamlFile(path)
	if err != nil {
		log.Error(err)
	}
	return err
}

func Config() *cfg.Config {
	return c
}
