package config

import (
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

// Config defines configuration for the antenna
type Config struct {
	HTTPAddress    string        `yaml:"http_address"`
	InfluxAddress  string        `yaml:"influx_addrss"`
	InfluxDatabase string        `yaml:"influx_database"`
	SyncTime       time.Duration `yaml:"sync_time"`
}

// Load provides loading of the config
func Load(path string) (*Config, error) {
	c := &Config{}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal([]byte(yamlFile), &c)
	if err != nil {
		return nil, fmt.Errorf("unable to ")
	}
	return c, nil
}
