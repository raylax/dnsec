package server

import (
	"github.com/miekg/dns"
	"github.com/raylax/dnsec/handler"
	"log"
)

type Server struct {
	Handler *handler.Handler
	srv *dns.Server
}

func (s *Server) Start() {
	s.srv = &dns.Server{Addr: ":53", Net: "udp", Handler: s.Handler}
	log.Println("Start DNS server")
	go func() {
		err := s.srv.ListenAndServe()
		if err != nil {
			log.Printf("Server error - %s\n", err.Error())
		}
	}()
}

func (s *Server) Shutdown() {
	if s.srv == nil {
		return
	}
	err := s.srv.Shutdown()
	if err != nil {
		log.Printf("Shutdown DNS server error - %s\n", err.Error())
	}
}