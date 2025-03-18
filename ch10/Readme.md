new things learned from each project:
* rServer
    * exercise 1: Include the functionality of binary.go in your own RESTful server
    * setup rest in main
        * Init mux routing with `mux := http.NewServeMux()`
        * assign each func to handler initiation with `mux.Handle("/endpoint", http.HandlerFunc(endpFunc))`
    * each handler func usual flows:
        * has 2 params (`(w http.ResponseWriter, r *http.Request`)
        * validate method first `r.method != http.MethodPost`
        * request body reading with `d := io.ReadAll(r.Body)`
        * extract json structure from body with `json.Unmarshal(d, &aStruct)`
        * reject invalid request with `r.WriteHeader(httpstatus)`

* rClient
    * on hit endpoint:
        * initiate http client (`c := &http.Client{}`)
        * form request body with `m := json.Marshal(aStruct)` and put in `u := bytes.NewReader(m)`
        * create request to http with `req := http.NewRequest(method, baseUrl+endpoint, u)`
        * execute request (`resp := c.Do(req)`)
        * read and parse response (`data := io.ReadAll(resp.Body)` -> `string(data)` -> `json.Unmarshal(data, &varStruct)` )

* restdb
    * using postgresdb with `database/sql` is pretty straighforward:
        * setup config in string (`conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Hostname, Port, Username, Password,Database)`)
        * init db conn with `sql.Open("postgresql", "host=")`
        * fetch data from db = `sql.Query` -> iterate `rows.Next` -> assign `rows.Scan(&var1, &var2, ...)` to var
        * mutate data to db = `stmt := sql.Prepare("UPDATE FROM table WHERE ID = $1")` -> `stmt.Exec(id)`

* binary
    * using `gorilla/mux` for routing with regex pattern matching
    * static files served with `http.FileServer`
    * binary file is brought along with `http.Request` body

* swagger
    * swagger generates docs based on comment in file
    * docs format for route handler:
    ```
    // swagger:route [HTTP_METHOD] [PATH] [TAG] [OPERATION_ID]
    // [Description of what the endpoint does]
    //
    // responses:
    //   [STATUS_CODE]: [RESPONSE_TYPE]
    //   [STATUS_CODE]: [RESPONSE_TYPE]
    ```
    * swagger models are defined with `// swagger:model` comment
    * swagger parameters are defined with `// swagger:parameters [OPERATION_ID]`
    * swagger responses are defined with `// swagger:response [RESPONSE_TYPE]`
    * in-line documentation for struct fields with `// [Description]` and annotations like `// in: body`, `// required: true`, etc.
    * `in: body` annotation indicates that the data will be sent in the JSON body of the HTTP request or returned in the response body.
    * swagger comment can refer to existing entity (`// 404: ErrorMessage`)
    * generating swagger documentation:
        * install required tools:
            * `go install github.com/go-swagger/go-swagger/cmd/swagger@latest` for spec generation
            * `brew install swagger-codegen` for HTML docs generation
        * generate swagger spec with `swagger generate spec -o ./swagger.yaml`
        * generate HTML docs with `swagger-codegen generate -i swagger.yaml -l html2 -o docs`
        * output will be in `docs/index.html` which can be viewed in browser
    * serving swagger docs in Go application:
        * use the `go-openapi/runtime/middleware` package
        * set up a GET subrouter: `getMux := mux.Methods(http.MethodGet).Subrouter()`
        * configure Redoc options: `opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}`
        * create handler: `sh := middleware.Redoc(opts, nil)`
        * register handlers:
            * `getMux.Handle("/docs", sh)` - serves the Redoc UI at /docs endpoint
            * `getMux.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))`- serves the swagger.yaml file
        * access the documentation UI at http://localhost:PORT/docs