package config

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func readConfig(in io.Reader, c map[string]interface{}) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(in)

	if err := yaml.Unmarshal(buf.Bytes(), &c); err != nil {
		fmt.Println("Error parsing config: ", err)
	}
}
