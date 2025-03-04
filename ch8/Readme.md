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