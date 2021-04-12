package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ResponseTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{Name: "external_url_response_ms"}, []string{"url"})
	Status       = prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: "external_url_up"}, []string{"url"})
)

func init() {
	prometheus.MustRegister(ResponseTime, Status)
}
