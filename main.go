package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("h", "", "Host")
	port := flag.Int("p", 0, "Port")
	flag.Parse()
	address := fmt.Sprintf("%s:%d", *host, *port)

	server := &RPCServer{}
	listener := HookListener{server: server}
	go listener.Listen()
	server.Serve(address)
}
