package config

import (
	"bytes"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
)

type Config struct {
	configPath string
	config     map[string]interface{}
	defaults   map[string]interface{}
}

func New() *Config {
	c := new(Config)
	c.configPath = "config.yml"
	c.config = make(map[string]interface{})
	c.defaults = make(map[string]interface{})
	return c
}

var c *Config

func init() {
	c = New()
}

func Read(path string) error {
	c.configPath = path
	fmt.Println("Reading configuration from", c.configPath)

	file, err := ioutil.ReadFile(c.configPath)
	if err != nil {
		fmt.Println("Error reading config:", err)
		return err
	}
	reader := bytes.NewReader(file)
	readConfig(reader, c.config)
	return nil
}

func (this *Config) Get(key string) interface{} {
	return this.config[key]
}

func GetString(key string) string {
	return c.GetString(key)
}
func (this *Config) GetString(key string) string {
	return cast.ToString(this.Get(key))
}

func GetInteger(key string) int {
	return c.GetInteger(key)
}
func (this *Config) GetInteger(key string) int {
	return cast.ToInt(this.Get(key))
}
