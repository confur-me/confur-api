// This code was initially taken from spf13/viper package by Steve Francia
// https://github.com/spf13/viper

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

// Returns the value associated with the key as a string
func GetString(key string) string {
	return c.GetString(key)
}
func (this *Config) GetString(key string) string {
	return cast.ToString(this.Get(key))
}

// Returns the value associated with the key as an integer
func GetInteger(key string) int {
	return c.GetInteger(key)
}
func (this *Config) GetInteger(key string) int {
	return cast.ToInt(this.Get(key))
}

// Returns the value associated with the key as a map of interfaces
func GetStringMap(key string) map[string]interface{} {
	return c.GetStringMap(key)
}
func (this *Config) GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(c.Get(key))
}

// Returns the value associated with the key as a map of strings
func GetStringMapString(key string) map[string]string {
	return c.GetStringMapString(key)
}
func (this *Config) GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(c.Get(key))
}

// Returns the value associated with the key asa boolean
func GetBool(key string) bool {
	return c.GetBool(key)
}
func (this *Config) GetBool(key string) bool {
	return cast.ToBool(c.Get(key))
}
