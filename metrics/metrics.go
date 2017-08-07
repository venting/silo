package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

//
var (

	// ContainersTotal - Count of containers on the system
	ContainersTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "silo_containers_total",
			Help: "Gauge of containers",
		}, []string{"state"})

	// ContainerEvents - Counter of container events
	ContainerEvents = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "silo_container_events",
			Help: "counter of container event operations",
		}, []string{"event_type", "status"})

	// Used for function duration tracking
	start = time.Now()
)

// Init registers the prometheus metrics for the measurement of the exporter itsself.
func Init() {

	prometheus.MustRegister(ContainersTotal)
	prometheus.MustRegister(ContainerEvents)

}
