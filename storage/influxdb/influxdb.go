package influxdb

import (
	"sync"

	influxdb "github.com/influxdb/influxdb/client"
	"github.com/saromanov/antenna/structs/v1"
	"github.com/saromanov/antenna/storage"
)

type influxDB struct {
	client   *influxdb.Client
	database string
	lock     sync.Mutex
}

func new(url string) (storage.Storage, error) {
	config := &influxdb.Config{
		URL:      url,
		Username: username,
		Password: password,
	}
	client, err := influxdb.NewClient(*config)
	if err != nil {
		return nil, err
	}
	return &influxDB{
		client: client,
	}, nil
}

// Add provides adding of stat
func (i*influxDB) Add(metrics *structs.ContainerStat) error {
	var points *[]influxdb.Point
	if err := c.Write(points); err != nil {
		return err
	}
	if err := c.Close(); err != nil {
    	return err
	}
	return nil
}
