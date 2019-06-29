package main

import (
	"crypto/tls"
	"flag"
	"net/http"
	"time"

	"github.com/saromanov/antenna/antenna"
	"github.com/saromanov/antenna/storage"
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
	st, err := influxdb.New(&storage.Config{
		URL: "http://localhost:8086",
	})
	if err != nil {
		log.WithFields(log.Fields{
			"stage": logStage,
		}).Fatalf("unable to init InfluxDB: %v", err)
	}

	log.WithFields(log.Fields{
		"stage": logStage,
	}).Info("init of HTTP client")

	client := createHTTPClient("", "")

	log.WithFields(log.Fields{
		"stage": logStage,
	}).Info("init of Antenna app")

	ant := antenna.Application{
		HTTPClient: client,
		Store:      st,
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
