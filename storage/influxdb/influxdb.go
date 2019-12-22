package influxdb

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"github.com/influxdata/influxdb-client-go"
	"github.com/pkg/errors"
	"github.com/saromanov/antenna/storage"
	structs "github.com/saromanov/antenna/structs/v1"
)

type influxDB struct {
	client   *influxdb.Client
	database string
	lock     sync.Mutex
}

// New creates storage based on name
// At the init stage, its supports only InfluxDB
func New(conf *storage.Config) (storage.Storage, error) {
	return new(conf)
}

func new(conf *storage.Config) (storage.Storage, error) {
	influx, err := influxdb.New(conf.URL, conf.Token, influxdb.WithUserAndPass(conf.Username, conf.Password))
	if err != nil {
		return nil, errors.Wrap(err, "unable to init influx client")
	}
	return &influxDB{
		client:   influx,
		database: conf.Database,
	}, nil
}

// Add provides adding of stat
func (i *influxDB) Add(metrics *structs.ContainerStat) error {
	i.lock.Lock()
	defer i.lock.Unlock()

	metricsConverted := i.toRowMetric(metrics)
	if _, err := i.client.Write(context.Background(), i.database, "test", metricsConverted...); err != nil {
		return errors.Wrap(err, "unable to write metrics")
	}
	return nil
}

// Search provides searching of the stats by the query
func (i *influxDB) Search(req *structs.ContainerStatSearch) ([]*structs.ContainerStat, error) {
	return nil, nil
}

// Search provides searching of the stats by the query
func (i *influxDB) Aggregate(req *structs.AggregateSearchRequest) (*structs.AggregateSearchResponse, error) {
	response, err := i.client.QueryCSV(context.TODO(), req.Request, i.database)
	if err != nil {
		return nil, errors.Wrap(err, "unable to query data")
	}
	aggr := &structs.AggregateSearchResponse{}
	b, err := ioutil.ReadAll(response)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(b))
	return aggr, nil
}

// Close provides closing of db instance
func (i *influxDB) Close() error {
	return i.client.Close()
}

func (i *influxDB) Info() map[string]interface{} {
	return map[string]interface{}{
		"name": "influxdb",
	}
}

func (i *influxDB) toRowMetric(metrics *structs.ContainerStat) []influxdb.Metric {
	points := []influxdb.Metric{}
	points = append(points, makeMetric("cache", metrics.Cache, metrics.Image))
	points = append(points, makeMetric("usage", metrics.Usage, metrics.Image))
	return points
}

// makePoints provides method for making point for InfluxDB
func makeMetric(name string, value interface{}, image string) *influxdb.RowMetric {
	return influxdb.NewRowMetric(
		map[string]interface{}{name: value},
		"antenna-metrics",
		map[string]string{"image": image},
		time.Now())
}
