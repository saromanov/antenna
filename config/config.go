package config

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/saromanov/antenna/storage"
	"gopkg.in/yaml.v2"
)

const (
	defaultAddress = "localhost:1255"
)

// Config defines configuration for the antenna
type Config struct {
	ServerAddress string          `yaml:"server_address"`
	SyncTime      time.Duration   `yaml:"sync_time"`
	Storage       *storage.Config `yaml:"storage"`
	Cert          string          `yaml:"cert_key"`
	Key           string          `yaml:"key"`
	IntervalCheck string          `yaml:"interval_check"`
}

// FillMissing provides filling of the missing mandatory values
func (c *Config) FillMissing() {
	if c.ServerAddress == "" {
		c.ServerAddress = defaultAddress
	}
	if c.IntervalCheck == "" {
		c.IntervalCheck = "@every 5s"
	}
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
		return nil, fmt.Errorf("unable to unmarshal config")
	}
	return c, nil
}

// LoadDefault provides loading of default data
func LoadDefault() *Config {
	return &Config{
		SyncTime:      15 * time.Second,
		ServerAddress: defaultAddress,
		Storage:       storage.LoadDefault(),
	}
}
