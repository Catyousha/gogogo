new things learned from each project:
* tcp-c
    * establish connection to tcp server with `c := net.Dial("tcp")`
    * write to dialed conn (`fmt.Fprintf(c, text+"\n")` / `c.Write(...)`) => send msg to server

* tcp-s
    * serving tcp server with `net.Listen("tcp")`
    * accept incoming request with `c := listener.Accept()`
    * write to listened conn (`c.Write(...)`) => send msg to connected client

* conc-tcp
    * exercise 3
    * tcp server can be served concurrently with goroutine

* udp-c
    * UDP address must be resolved through `s := net.ResolveUDPAddr("udp4", addr)` before dialed with `net.DialUDP("udp4", nil, s)`

* udp-s
    * UDP address must be resolved through `s := net.ResolveUDPAddr("udp4", addr)` before open to listen with `conn := net.ListenUDP("udp4", s)`
    * read message by dialed client with `conn.ReadFromUDP(assignedVar)`
    * write msg to client with `conn.WriteToUDP(text)`

* socket-client & socket-server
    * UNIX domain socket is used for establish communication in same machine.
    * listen & dial through `.socket` file

* ws-s
    * package `gorilla/websocket` can be used to implement websocket server
    * initiate websocket by upgrading http writer with `ws := websocket.Upgrader{}.Upgrader(httpWriter, httpReader, nil)`
    * get msg from client with `msgType, msgContent := ws.ReadMessage()`
    * write msg to client with `ws.WriteMessage(msgType, text)`

* ws-c
    * establish websocket connection through `websocket.DefaultDialer.Dial(url, nil)`
    * read message from server with `ws.ReadMessage()`
    * write message to server with `ws.WriteMessage(websocket.TextMessage, text)`
    * handle interrupt signal with `os.Signal` channel and `signal.Notify()`
    * implement timeout with `time.After()` and `syscall.Kill()`

* exercises 1, 2, 4, 5