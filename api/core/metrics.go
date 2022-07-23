package core

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var RequestsProcessedMetric = promauto.NewCounter(prometheus.CounterOpts{
	Name: "homies_requests_processed",
	Help: "The total number of processed requests",
})
