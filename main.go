package antenna

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/saromanov/antenna/antenna"
	"github.com/saromanov/antenna/storage"
	"github.com/saromanov/antenna/storage/influxdb"
)

var (
	port               = flag.Int("port", 8080, "port")
	prometheusEndpoint = flag.String("prometheus_endpoint", "/metrics", "Endpoint for export metrics")
)

func main() {
	st := influxdb.New(&storage.Config{
		URL: "//",
	})

	client := createHTTPClient()

	ant := antenna.Application{
		client: client,
	}
}

func createHTTPClient(cert, key string) http.Client {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	if collectorCert != "" {
		certData, err := tls.LoadX509KeyPair(collectorCert, collectorKey)
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
	return &http.Client{
		Transport: tr,
		Timeout:   time.Second * 10,
	}
}
