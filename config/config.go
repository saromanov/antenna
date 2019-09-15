package config

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/saromanov/antenna/storage"
	"gopkg.in/yaml.v2"
)

// Config defines configuration for the antenna
type Config struct {
	ServerAddress  string          `yaml:"server_address"`
	InfluxAddress  string          `yaml:"influx_addrss"`
	InfluxDatabase string          `yaml:"influx_database"`
	SyncTime       time.Duration   `yaml:"sync_time"`
	Storage        *storage.Config `yaml:"storage"`
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

func LoadDefault() *Config {
	return &Config{
		InfluxAddress:  "http://localhost:8086",
		InfluxDatabase: "antenna_container_metrics",
		SyncTime:       15 * time.Second,
		ServerAddress:  "localhost:1255",
		Storage: &storage.Config{
			URL:      "http://localhost:8086",
			Database: "antenna_container_metrics",
		},
	}
}
