package antenna

import (
	"crypto/tls"
	"flag"
	"net/http"
	"time"

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
}

func createHTTPClient() http.Client {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
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
