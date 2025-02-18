new things learned from each project:
* variadic
    * accept multiple args => (`...args float64`)
    * spread slice to args => `funcWithMultArgs(...slice)`

* defer
    * defer is Last In First Out (LIFO). it would be executed after all non-defered syntaxes inside func body is executed first.

* postGo
    * `docker-compose.yml` has basic config to setup postgres
    * package can be stored in github

* gogogo-gh-action (https://github.com/Catyousha/gogogo-gh-action)
    * `Dockerfile` to build and execute project
    * `.github/workflows/main.yml` to setup runner and execute docker
        * there are useful presets for runner (`actions/checkout@v2` & `actions/setup-go@v2`)
    * meanwhile gitlab runner is configured at `gitlab-ci.yml`

* getSchema
    * exercise 2
    * fetch db flow: open conn, query data and assign result to variable through reference, iterate each rows, assign each data to variable through reference again.