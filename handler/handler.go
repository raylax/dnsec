package handler

import (
	"github.com/miekg/dns"
	"github.com/raylax/dnsec/client"
	"log"
	"net"
	"strings"
	"time"
)

type Handler struct {
	Client client.Client
}


func (s *Handler) ServeDNS(w dns.ResponseWriter, q *dns.Msg) {
	start := time.Now()
	r, err := s.Client.Query(q)
	remote := w.RemoteAddr().(*net.UDPAddr)
	if err != nil {
		log.Printf("[%s] Query error - %s\n", remote.IP.String(), err.Error())
		return
	}
	since := time.Since(start)
	log.Printf("[%s] [%dms] %s", remote.IP.String(), since.Milliseconds(), getNames(q))
	w.WriteMsg(r)
}

func getNames(q *dns.Msg) string {
	if len(q.Question) == 0 {
		return ""
	}
	names := make([]string, len(q.Question))
	for i, question := range q.Question {
		names[i] = dns.Type(question.Qtype).String() + " " + question.Name
	}
	return strings.Join(names, ", ")
}