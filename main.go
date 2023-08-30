package main

import (
	"TCC2/server"
	"flag"
	"log"
)

func main() {
	var addr, port, url string
	flag.StringVar(&addr, "addr", "127.0.0.1", "server listen ip address")
	flag.StringVar(&port, "port", "8080", "server tcp port")
	flag.StringVar(&url, "chain", "", "rest api to concatenate request.")
	flag.Parse()
	sv := server.SetupServer(url)
	if sv == nil {
		return
	}
	err := sv.Run(addr + ":" + port)
	log.Fatal(err)
}
