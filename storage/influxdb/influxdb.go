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
	client       *influxdb.Client
	database     string
	organization string
	lock         sync.Mutex
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
		client:       influx,
		database:     conf.Database,
		organization: "test",
	}, nil
}

// Add provides adding of stat
func (i *influxDB) Add(metrics *structs.ContainerStat) error {
	i.lock.Lock()
	defer i.lock.Unlock()

	metricsConverted := toRowMetric(metrics)
	if _, err := i.client.Write(context.Background(), i.database, i.organization, metricsConverted...); err != nil {
		fmt.Println("ERRR: ", err)
		return errors.Wrap(err, "unable to write metrics")
	}
	return nil
}

// Search provides searching of the stats by the query
func (i *influxDB) Search(req *structs.ContainerStatSearch) ([]*structs.ContainerStat, error) {
	if req == nil || req.Request == "" {
		return nil, fmt.Errorf("request is not defined")
	}
	response, err := i.client.QueryCSV(context.TODO(), req.Request, i.database)
	if err != nil {
		return nil, errors.Wrap(err, "unable to query data")
	}
	resp := []*structs.ContainerStat{}
	for response.Next() {
		var stat structs.ContainerStat
		if err := response.Unmarshal(&stat); err != nil {
			return nil, err
		}
		resp = append(resp, &stat)
	}
	return resp, nil
}

// Search provides searching of the stats by the query
func (i *influxDB) Aggregate(req *structs.AggregateSearchRequest) (*structs.AggregateSearchResponse, error) {
	if req == nil || req.Request == "" {
		return nil, fmt.Errorf("request is not defined")
	}
	response, err := i.client.QueryCSV(context.TODO(), req.Request, i.database)
	if err != nil {
		return nil, errors.Wrap(err, "unable to query data")
	}

	for response.Next() {
		var stat structs.ContainerStat
		if err := response.Unmarshal(&stat); err != nil {
			return nil, err
		}
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

// Info returns information about storage
func (i *influxDB) Info() map[string]interface{} {
	msg := "OK"
	if err := i.client.Ping(context.TODO()); err != nil {
		msg = err.Error()
	}

	return map[string]interface{}{
		"name":                   "influxdb",
		"client":                 msg,
		"metrics_avaailable_num": 11,
	}
}

// Note: Influxdb don't understand uint types. Its converting to int
func toRowMetric(metrics *structs.ContainerStat) []influxdb.Metric {
	points := []influxdb.Metric{}
	points = append(points, makeMetric("cache", metrics.Cache, metrics.Image, metrics.Name))
	points = append(points, makeMetric("usage", metrics.Usage, metrics.Image, metrics.Name))
	points = append(points, makeMetric("cpu_total_usage", metrics.CPU.TotalUsage, metrics.Image, metrics.Name))
	points = append(points, makeMetric("cpu_online", metrics.CPU.OnlineCPUs, metrics.Image, metrics.Name))
	points = append(points, makeMetric("num_procs", metrics.NumProcs, metrics.Image, metrics.Name))
	points = append(points, makeMetric("read_size_bytes", metrics.ReadSizeBytes, metrics.Image, metrics.Name))
	points = append(points, makeMetric("write_size_bytes", metrics.WriteSizeBytes, metrics.Image, metrics.Name))
	points = append(points, makeMetric("tx_packets", metrics.TxPackets, metrics.Image, metrics.Name))
	points = append(points, makeMetric("tx_dropped", metrics.TxDropped, metrics.Image, metrics.Name))
	points = append(points, makeMetric("tx_errors", metrics.TxErrors, metrics.Image, metrics.Name))
	points = append(points, makeMetric("tx_bytes", metrics.TxBytes, metrics.Image, metrics.Name))
	return points
}

// makePoints provides method for making point for InfluxDB
func makeMetric(name string, value interface{}, image, containerName string) *influxdb.RowMetric {
	return influxdb.NewRowMetric(
		map[string]interface{}{name: value},
		"antenna-metrics",
		map[string]string{"image": image, "name": containerName},
		time.Now().UTC())
}
