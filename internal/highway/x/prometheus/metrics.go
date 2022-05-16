package prometheus

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sonr-io/sonr/pkg/host"
)

type HighwayTelemetry struct {
	endpoint string
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "Sonr",
		Name:      "total events processed",
		Help:      "The total number of processed events on the highway",
	})

	objectsAdded = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "Sonr",
		Subsystem: "data storage",
		Help:      "Counts the number of objects added to the highway",
	})

	blobsAdded = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "Sonr",
		Subsystem: "file storage",
		Help:      "Counts the number of objects added to the highway",
	})
)

func New(ctx context.Context, hn host.SonrHost) (*HighwayTelemetry, error) {
	defer RegisterEvents()
	return &HighwayTelemetry{endpoint: "/metrics"}, nil
}

func (ht *HighwayTelemetry) RecordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

func (ht *HighwayTelemetry) GetMetricsHandler() http.Handler {
	defer ht.RecordMetrics()
	return promhttp.Handler()
}
