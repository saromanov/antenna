package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metricValue struct {
	value  float64
	labels []string
}

type containerMetric struct {
	name        string
	help        string
	valueType   prometheus.ValueType
	extraLabels []string
}

type PrometheusCollector struct {
	infoProvider     infoProvider
	errors           prometheus.Gauge
	containerMetrics []containerMetric
}
