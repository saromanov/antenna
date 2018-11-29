package influxdb

import (
	"sync"
	"fmt"
	influxdb "github.com/influxdata/influxdb/client/v2"
	"github.com/saromanov/antenna/structs/v1"
	"github.com/saromanov/antenna/storage"
)

type influxDB struct {
	client   *influxdb.Client
	database string
	lock     sync.Mutex
}

// New creates storage based on name
// At the init stage, its supports only InfluxDB
func New(conf *Config) (Storage, error) {
	return new(conf.URL)
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
	i.lock.Lock()
	defer i.lock.Unlock()
	var points []*influxdb.Point
	if err := c.Write(points); err != nil {
		return err
	}
	if err := c.Close(); err != nil {
    	return err
	}
	return nil
}

func (i*influxDB) toPoints(metrics *structs.ContainerStat)[]*influxdb.Point {
	if len(metrics) == 0 {
		return nil
	}
	points := []*influxdb.Point{}
	points = append(points, makePoint("cpu", metrics.CPU)
	return nil
}

// makePoints provides method for making point for InfluxDB
func makePoint(name string, value interface{}) *influxdb.Point {
	fields := map[string]interface{}{
		fieldValue: value,
	}

	return &influxdb.Point{
		Measurement: name,
		Fields:      fields,
	}
}
