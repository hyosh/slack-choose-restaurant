package main

import (
	"flag"
	"math/rand"
	"time"
)

const Version = "1.0.0"

func main() {
	port := flag.Int("p", 8080, "port to listen on")
	host := flag.String("h", "", "host to listen on")
	versionFlag := flag.Bool("v", false, "print version and exit")

	flag.Parse()

	if *versionFlag {
		println("Version:", Version)
		return
	}

	// init random
	rand.Seed(time.Now().UnixNano())

	server := NewServer(*port, *host)
	server.Start()

}
