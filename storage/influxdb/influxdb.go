package influxdb

import (
	"sync"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/saromanov/antenna/storage"
	structs "github.com/saromanov/antenna/structs/v1"
)

type influxDB struct {
	client   client.Client
	database string
	lock     sync.Mutex
}

// New creates storage based on name
// At the init stage, its supports only InfluxDB
func New(conf *storage.Config) (storage.Storage, error) {
	return new(conf)
}

func new(conf *storage.Config) (storage.Storage, error) {
	config := client.HTTPConfig{
		Addr:     conf.URL,
		Username: conf.Username,
		Password: conf.Password,
	}
	cli, err := client.NewHTTPClient(config)
	if err != nil {
		return nil, err
	}
	return &influxDB{
		client: cli,
	}, nil
}

// Add provides adding of stat
func (i *influxDB) Add(metrics *structs.ContainerStat) error {
	i.lock.Lock()
	defer i.lock.Unlock()
	var points []*client.Point
	if err := i.client.Write(points); err != nil {
		return err
	}
	if err := i.client.Close(); err != nil {
		return err
	}
	return nil
}

// Close provides closing of db instance
func (i *influxDB) Close() error {
	i.client.Close()
}

func (i *influxDB) toPoints(metrics *structs.ContainerStat) []*client.Point {
	if len(metrics) == 0 {
		return nil
	}
	points := []*client.Point{}
	points = append(points, makePoint("cpu", metrics.CPU))
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
