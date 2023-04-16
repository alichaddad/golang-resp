package main

import (
	"flag"
	"os"

	"log"

	"github.com/alichaddad/goresp/server"
)

func main() {
	var address string
	flag.StringVar(&address, "address", "localhost:7000", "server address")
	flag.Parse()
	server := server.New(address)
	err := server.Start()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
