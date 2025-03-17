new things learned from each project:
* rServer
    * setup rest in main
        * Init mux routing with `mux := http.NewServeMux()`
        * assign each func to handler initiation with `mux.Handle("/endpoint", http.HandlerFunc(endpFunc))`
    * each handler func usual flows:
        * has 2 params (`(w http.ResponseWriter, r *http.Request`)
        * validate method first `r.method != http.MethodPost`
        * request body reading with `d := io.ReadAll(r.Body)`
        * extract json structure from body with `json.Unmarshal(d, &aStruct)`
        * reject invalid request with `r.WriteHeader(httpstatus)`