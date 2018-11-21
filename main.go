package antenna

import (
	"flag"

	"github.com/saromanov/antenna/storage"
)

var (
	port               = flag.Int("port", 8080, "port")
	prometheusEndpoint = flag.String("prometheus_endpoint", "/metrics", "Endpoint for export metrics")
)

func main() {
	st := storage.New(&storage.Config{
		Name: "influxdb",
	})
}
