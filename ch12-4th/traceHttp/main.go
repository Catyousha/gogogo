package main

import (
	"fmt"
	"net/http"
	"net/http/httptrace"
	"os"
)



func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: URL\n")
		return
	}

	URL := os.Args[1]
	client := http.Client{}

	// setup tracer
	req, _ := http.NewRequest("GET", URL, nil)
	trace := &httptrace.ClientTrace{
		GotFirstResponseByte: func() {
			fmt.Println("First response byte!")
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("Dial start")
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("Dial done")
		},
		WroteHeaders: func() {
			fmt.Println("Wrote headers")
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	
	// tracing
	fmt.Println("Requesting data from server!")
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// go run . http://www.google.com
	// Requesting data from server!
	// DNS Info: {Addrs:[{IP:xxx.xxx.xx.xxx Zone:} {IP:xxxx:xxxx:xxxx:xx::xx Zone:}] Err:<nil> Coalesced:false}
	// Dial start
	// Dial done
	// Got Conn: {Conn:0x14000058128 Reused:false WasIdle:false IdleTime:0s}
	// Wrote headers
	// First response byte!
	// DNS Info: {Addrs:[{IP:xxx.xxx.xx.xxx Zone:} {IP:xxxx:xxxx:xxxx:xx::xx Zone:} {IP:xxxx:xxxx:xxxx:xx::xx Zone:} {IP:xxxx:xxxx:xxxx:xx::xx Zone:} {IP:xxxx:xxxx:xxxx:xx::xx Zone:} {IP:xxxx:xxxx:xxxx:xx::xx Zone:}] Err:<nil> Coalesced:false}
	// Dial start
	// Dial done
	// Got Conn: {Conn:0x1400019e000 Reused:false WasIdle:false IdleTime:0s}
	// Wrote headers
	// First response byte!
}
