package main

import (
	"flag"
	"math/rand"
	"time"
)

func main() {
	port := flag.Int("p", 8080, "port to listen on")
	host := flag.String("h", "", "host to listen on")

	flag.Parse()

	// init random
	rand.Seed(time.Now().UnixNano())

	server := NewServer(*port, *host)
	server.Start()

}
