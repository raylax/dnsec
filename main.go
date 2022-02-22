package main

import (
	"fmt"
	"github.com/raylax/dnsec/client"
	"github.com/raylax/dnsec/handler"
	"github.com/raylax/dnsec/server"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)


func printUsage() {
	fmt.Printf("Usage %s [ENDPOINT]\n", os.Args[0])
	os.Exit(0)
}

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM)
	signal.Notify(signalChan, syscall.SIGINT)
	args := os.Args[1:]
	if len(args) < 1 {
		printUsage()
	}
	endpoint := args[0]
	var c client.Client
	if strings.HasPrefix(endpoint, "http") {
		c = client.NewDoH(endpoint)
		log.Println("[*] Using DoH")
	} else {
		c = client.NewDoT(endpoint)
		log.Println("[*] Using DoT")
	}
	log.Printf("[*] Endpoint [%s]\n", endpoint)
	h := &handler.Handler{Client: c}
	s := &server.Server{Handler: h}
	s.Start()

	<-signalChan
	log.Println("Shutdown")
	s.Shutdown()
}
