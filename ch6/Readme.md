new things learned from each project:

* signals
    * user can sent command to program which interpreted as signal
    * preserve `os.Signal` memory (`make(os.Signal, 1)`) to store user signal as variable
    * `signals.Notify(signalVar)` to bind incoming signal to var
    * goroutine to run infinite non-blocking commands
    * `awaitingVar := <- incomingDataVar`

* ioInterface
    * terminal reader can be replaced by `bufio.NewReader(&readerThatImplementRead)`
    * result value can be assigned without specify its return
        * `func (s *S2) Read(p []byte) (n int, err error)`
            * `err = io.EOF`
            * `n++`

* encodeDecode
    * json object has struct type with this format:
        * `Field string \`json:"json_field"\``
    * encode to str with `json.Marshal(&obj)`
    * decode from str with `json.Unmarshall([]byte(str), &targetObj)`

* phonebook-v5
    * simple crud app with cobra cli
    * each files in cmd subpackage has `init()` func, auto running when `cmd.Execute()` is called from `root.go`.
