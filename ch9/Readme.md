new things learned from each project:
* tcp-c
    * establish connection to tcp server with `c := net.Dial("tcp")`
    * write to dialed conn (`fmt.Fprintf(c, text+"\n")` / `c.Write(...)`) => send msg to server

* tcp-s
    * serving tcp server with `net.Listen("tcp")`
    * accept incoming request with `c := listener.Accept()`
    * write to listened conn (`c.Write(...)`) => send msg to connected client

* conc-tcp
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