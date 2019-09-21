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
	ServerAddress string          `yaml:"server_address"`
	SyncTime      time.Duration   `yaml:"sync_time"`
	Storage       *storage.Config `yaml:"storage"`
	Cert          string          `yaml:"cert_key"`
	Key           string          `yaml:"key"`
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
		SyncTime:      15 * time.Second,
		ServerAddress: "localhost:1255",
		Storage: &storage.Config{
			URL:      "http://localhost:9999",
			Database: "antenna_metrics_data",
			Username: "test",
			Password: "12345678",
			Token:    "ScOgVIc-Ti6OCTjx4CCfUwiQVm3XvA0cPZEh-NIE0YbXd81izGQ_0Y8f9ZaSw89MtHQv3QS0pkhqCLcm5ju_4w==",
		},
	}
}
