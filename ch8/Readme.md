new things learned from each project:
* wwwServer
    * describe route in handler, bind handler with `http.HandleFunc('/route', handler)`
    * serve http server with `http.ListenAndServe(PORT, nil)`
    * html response can be printed out by `fmt.Printf("<tag>body</tag>")`

* www-phone
    * route can include parameters after `/`, e.g., `/insert/john/doe/123-456`
    * parameters can be extracted using `strings.Split(r.URL.Path, "/")`
    * store data in memory with struct and slice
    * persist data using CSV file
    * use map as index for fast lookup
    * separate handlers into different functions for better organization
    * input validation with regex using `regexp.MustCompile`
    * http request workflow: mux routes request to handler -> handler extracts parameters -> validates input -> performs operation (search/insert/delete) on in-memory data -> updates CSV file -> returns response

* metrics
    * tracks number of active goroutines using runtime metrics
    * uses `metrics.Sample` to hold metric data
    * uses `metrics.Read()` to get current metric values (how many processes are running)

* sample-pro
    * register metrics with `prometheus.MustRegister()`
    * counter metrics continuously increase with `counter.Add()`
    * gauge metrics can fluctuate up/down with `gauge.Add()`
    * histogram metrics show value distribution in buckets with `histogram.Observe()`
    * summary metrics track count and sum of observations with `summary.Observe()`
    * expose metrics endpoint with `http.Handle("/metrics", promhttp.Handler())`
    * metrics can be viewed using curl on /metrics endpoint

* prometheus
    * basic concept of integration between docker and prometheus + grafana
    * collect runtime metrics using `metrics.Read()` with specific metric paths
    * use garbage collection with `runtime.GC()` to manage memory
    * store metrics in `metrics.sample`, then snapshot it into prometheus with `prometheus.Gauge.Set( float64(getMetric[0].Value.Uint64()) )`