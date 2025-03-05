package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var PORT = ":1234"

var counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Namespace: "ctysh",
		Name:      "my_counter",
		Help:      "This is my counter",
	})

var gauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "ctysh",
		Name:      "my_gauge",
		Help:      "This is my gauge",
	})

var histogram = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Namespace: "ctysh",
		Name:      "my_histogram",
		Help:      "This is my histogram",
	})

var summary = prometheus.NewSummary(
	prometheus.SummaryOpts{
		Namespace: "ctysh",
		Name:      "my_summary",
		Help:      "This is my summary",
	})

func main() {
	rand.NewSource(42)

	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)

    // Counter continuously increases
    // Gauge fluctuates
    // Histogram shows distribution in predefined buckets
    // Summary shows total count and sum of observations
    go func() {
        // each two seconds,
        // counter: add random val between 0-5
        // gauge: add random val between -5 to +10
        // histogram: observe random val between 0-10
        // summary: observe random val between 0-10
        for {
            counter.Add(rand.Float64() * 5)
            gauge.Add(rand.Float64()*15 - 5)
            histogram.Observe(rand.Float64() * 10)
            summary.Observe(rand.Float64() * 10)
            time.Sleep(2 * time.Second)
        }
    }()

    http.Handle("/metrics", promhttp.Handler())
    fmt.Println("Server started at", PORT)
    fmt.Println(http.ListenAndServe(PORT, nil))

    // curl localhost:1234/metrics | grep ctysh
    /// Counter
    // # HELP ctysh_my_counter This is my counter
    // # TYPE ctysh_my_counter counter
    // ctysh_my_counter 24.35907131561028

    /// Gauge
    // # HELP ctysh_my_gauge This is my gauge
    // # TYPE ctysh_my_gauge gauge
    // ctysh_my_gauge 35.78831259375473

    /// Histogram
    // # HELP ctysh_my_histogram This is my histogram
    // # TYPE ctysh_my_histogram histogram
    // ctysh_my_histogram_bucket{le="0.005"} 0
    // ctysh_my_histogram_bucket{le="0.01"} 0
    // ctysh_my_histogram_bucket{le="0.025"} 0
    // ctysh_my_histogram_bucket{le="0.05"} 0
    // ctysh_my_histogram_bucket{le="0.1"} 0
    // ctysh_my_histogram_bucket{le="0.25"} 0
    // ctysh_my_histogram_bucket{le="0.5"} 0
    // ctysh_my_histogram_bucket{le="1"} 1
    // ctysh_my_histogram_bucket{le="2.5"} 4
    // ctysh_my_histogram_bucket{le="5"} 6
    // ctysh_my_histogram_bucket{le="10"} 9
    // ctysh_my_histogram_bucket{le="+Inf"} 9
    // ctysh_my_histogram_sum 37.23989460831672
    // ctysh_my_histogram_count 9

    /// Summary
    // # HELP ctysh_my_summary This is my summary
    // # TYPE ctysh_my_summary summary
    // ctysh_my_summary_sum 50.9547609146948
    // ctysh_my_summary_count 9
}
