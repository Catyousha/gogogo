new things learned from each project:
    * profileCla
        * `pprof` can be used to profiling cpu usage & memory
        * open profiled result with `go tool pprof <file>`
    * profileHttp
        * `pprof` can be served as handler, accessed through debug route
    
    * traceHttp
        * trace packet trip with `"net/http/httptrace"` package
    
    * table
        * present example of table-driven test
        * declare test func by using `(t *testing.T)` param
        * paralellize tests with `t.Parallel()`
    
    * testHttp
        * exercise 2: The code in testHTTP/server_test.go uses the same value for lastlogin in the expected variable. This is clearly a bug in restdb.go as the value of lastlogin should be updated. After correcting the bug, modify testHTTP/server_test.go to take into account the different values of the lastlogin field.
        * testing http flow:
            * initiate request (`http.NewRequest("METHOD", "/path", nil)`)
                -> setup recorder (`rr := httptest.NewRecorder()`)
                -> set handler func to mini http server (`handler := http.HandlerFunc(TimeHandler)`)
                -> bind recorder to request (`handler.ServeHTTP(rr, req)`)
                -> compare actual result from recorder to expected res (`if rr.Code != http.StatusOK: ...`)