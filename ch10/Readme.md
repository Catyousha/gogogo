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