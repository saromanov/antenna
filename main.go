package main

import (
	"crypto/tls"
	"flag"
	"net/http"
	"time"

	"github.com/saromanov/antenna/antenna"
	"github.com/saromanov/antenna/config"
	"github.com/saromanov/antenna/server"
	"github.com/saromanov/antenna/storage"
	"github.com/saromanov/antenna/storage/hashmap"
	"github.com/saromanov/antenna/storage/influxdb"
	log "github.com/sirupsen/logrus"
)

var (
	port               = flag.Int("port", 8080, "port")
	prometheusEndpoint = flag.String("prometheus_endpoint", "/metrics", "Endpoint for export metrics")

	logStage = "init"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	conf, err := config.Load("config.yaml")
	if err != nil {
		log.WithFields(log.Fields{
			"stage": logStage,
		}).Warnf("unable to load config: %v", err)
		conf = config.LoadDefault()
	}
	st, err := influxdb.New(&storage.Config{
		URL:      conf.InfluxAddress,
		Database: conf.InfluxDatabase,
	})
	if err != nil {
		log.WithFields(log.Fields{
			"stage": logStage,
		}).Fatalf("unable to init InfluxDB: %v", err)
	}

	log.WithFields(log.Fields{
		"stage": logStage,
	}).Infof("init of the server at the address %s", conf.ServerAddress)

	go server.Start(st, conf.ServerAddress)

	log.WithFields(log.Fields{
		"stage": logStage,
	}).Info("init of HTTP client")

	client := createHTTPClient("", "")

	log.WithFields(log.Fields{
		"stage": logStage,
	}).Info("init of Antenna app")

	hash, err := hashmap.New(&storage.Config{})
	if err != nil {
		log.WithFields(log.Fields{
			"stage": logStage,
		}).Fatalf("unable to init hashmap storage: %v", err)
	}
	ant := antenna.Application{
		HTTPClient: client,
		Store:      st,
		MapStore:   hash,
	}
	ant.Start()
}
func createHTTPClient(cert, key string) http.Client {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	if cert != "" {
		certData, err := tls.LoadX509KeyPair(cert, key)
		if err != nil {
			log.Fatalf("Failed to use certs %v", err)
		}

		tlsConfig.Certificates = []tls.Certificate{certData}
		tlsConfig.BuildNameToCertificate()
	}

	tr := &http.Transport{
		IdleConnTimeout: 30 * time.Second,
		TLSClientConfig: tlsConfig,
	}
	return http.Client{
		Transport: tr,
		Timeout:   time.Second * 10,
	}
}
