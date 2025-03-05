package main

import (
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"runtime/metrics"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var PORT = ":1234"

var n_goroutines = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "ctysh",
		Name:      "n_goroutines",
		Help:      "Number of goroutines running",
	})

var n_memory = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "ctysh",
		Name:      "n_memory",
		Help:      "Memory usage",
	})

func main() {
	rand.NewSource(time.Now().UnixNano())

	prometheus.MustRegister(n_goroutines)
	prometheus.MustRegister(n_memory)

	// The number of goroutines currently running.
	const nGo = "/sched/goroutines:goroutines"

	// The amount of heap memory in use.
	const nMem = "/memory/classes/heap/free:bytes"

	getMetric := make([]metrics.Sample, 2)
	getMetric[0].Name = nGo
	getMetric[1].Name = nMem

	// runs indefinitely, periodically collecting runtime metrics and updating the Prometheus gauges.
	go func() {
		for {
			// three new goroutines are spawned every iteration.
			for i := 1; i < 4; i++ {
				go func() {
					// allocates a large slice ([]int with 1 million elements) to simulate memory usage.
					_ = make([]int, 1_000_000)

					// sleeps for a random duration (0–9 seconds) before exiting
					time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
				}()
			}

			runtime.GC()

			// reads the current values of the runtime metrics into the getMetric slice.
			metrics.Read(getMetric)

			// put metrics to the Prometheus gauges.
			goVal := getMetric[0].Value.Uint64()
			memVal := getMetric[1].Value.Uint64()
			time.Sleep(time.Duration(rand.Intn(15)) * time.Second)
			
			n_goroutines.Set(float64(goVal))
			n_memory.Set(float64(memVal))
		}
	}()
	
	log.Println("Listening to port", PORT)
	http.Handle("/metrics", promhttp.Handler())
	log.Println(http.ListenAndServe(PORT, nil))
	

	// # HELP ctysh_n_goroutines Number of goroutines running
	// # TYPE ctysh_n_goroutines gauge
	// ctysh_n_goroutines 5
	
	// # HELP ctysh_n_memory Memory usage
	// # TYPE ctysh_n_memory gauge
	// ctysh_n_memory 139264
}
