package main

import (
	"context"
	"fmt"
	"github.com/raylax/DoH-local/client"
	"github.com/raylax/DoH-local/server"
	"log"
	"os"
)

var endpoint = "https://208.67.222.222/dns-query"
var listen = "0.0.0.0"

func printUsage() {
	fmt.Printf("Usage %s [ENDPOINT] [LISTEN]\n", os.Args[0])
	os.Exit(0)
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		printUsage()
	}
	if len(args) > 0 {
		endpoint = args[0]
	}
	if len(args) > 1 {
		listen = args[1]
	}
	log.Printf("[*] Using endpoint [%s]\n", endpoint)
	log.Printf("[*] Listen on [%s]\n", listen)
	c := client.Client{
		Endpoint: endpoint,
	}
	s := server.Server{
		Ctx:    context.Background(),
		Client: c,
		Listen: "0.0.0.0",
	}
	if err := s.Start(); err != nil {
		log.Fatalln("[!] ERR - " + err.Error())
	}
}
