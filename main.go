package antenna

import (
	"flag"

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
