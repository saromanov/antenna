package antenna

import (
	"flag"
)

var (
	port               = flag.Int("port", 8080, "port")
	prometheusEndpoint = flag.String("prometheus_endpoint", "/metrics", "Endpoint for export metrics")
)

func main() {

}
