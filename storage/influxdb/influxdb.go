package influxdb

import (
	"sync"

	"github.com/influxdata/influxdb/client/v2"
	"github.com/pkg/errors"
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
		database: conf.Database,
	}, nil
}

// Add provides adding of stat
func (i *influxDB) Add(metrics *structs.ContainerStat) error {
	i.lock.Lock()
	defer i.lock.Unlock()
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.database,
		Precision: "s",
	})
	if err != nil {
		return errors.Wrap(err, "unable to create new batch point")
	}
	if err := i.client.Write(bp); err != nil {
		return errors.Wrap(err, "unable to write data")
	}
	if err := i.client.Close(); err != nil {
		return errors.Wrap(err, "unable to close connection")
	}
	return nil
}

// Search provides searching of the stats by the query
func (i *influxDB) Search(req *structs.ContainerStatSearch) ([]*structs.ContainerStat, error) {
	return nil, nil
}

// Close provides closing of db instance
func (i *influxDB) Close() error {
	return i.client.Close()
}

func (i *influxDB) toPoints(metrics *structs.ContainerStat) []*client.Point {
	points := []*client.Point{}
	points = append(points, makePoint("cpu", "1"))
	return nil
}

// makePoints provides method for making point for InfluxDB
func makePoint(name string, value interface{}) *client.Point {
	return nil
}
