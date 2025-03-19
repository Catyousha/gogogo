new things learned from each project:
    * profileCla
        * `pprof` can be used to profiling cpu usage & memory
        * open profiled result with `go tool pprof <file>`
    * profileHttp
        * `pprof` can be served as handler, accessed through debug route
    
    * traceHttp
        * trace packet trip with `"net/http/httptrace"` package
        * 