package storage

import (
	"os"
	structs "github.com/saromanov/antenna/structs/v1"
)

// Storage defines main interface for Providing
// handling of storage data
type Storage interface {
	Add(*structs.ContainerStat) error
	Close() error
	Search(*structs.ContainerStatSearch) ([]*structs.ContainerStat, error)
	Aggregate(*structs.AggregateSearchRequest) (*structs.AggregateSearchResponse, error)
	Info() map[string]interface{}
}

// Config defines configuration for Storage init
type Config struct {
	Organization string
	Name         string
	URL          string
	Username     string
	Password     string
	Database     string
	Token        string
}

// LoadDefault provides loading of default storage
func LoadDefault() *Config {
	return &Config{
		URL:      os.Getenv("ANTENNA_STORAGE_URL"),
		Database: os.Getenv("ANTENNA_STORAGE_DATABASE"),
		Username: os.Getenv("ANTENNA_STORAGE_USERNAME"),
		Password: os.Getenv("ANTENNA_STORAGE_PASSWORD"),
		Token:    os.Getenv("ANTENNA_STORAGE_TOKEN"),
	}
}
