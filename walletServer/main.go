package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("Wallet Server:")
}

func main() {
	port := flag.Uint("port", 3000, "TCP port for wallet server.")
	gateway := flag.String("g", "http://127.0.0.1:5000", "TCP gateway for host server")
	flag.Parse()
	ws := NewWalletServer(uint16(*port), *gateway)
	ws.Run()
}
