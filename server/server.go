package server

import (
	"context"
	"github.com/miekg/dns"
	"github.com/raylax/DoH-local/client"
	"log"
	"net"
)

const bufferSize = 1024

type Server struct {
	Ctx context.Context
	Client client.Client
	Listen string
}

func (s Server) Start() error {
	pc, err := net.ListenPacket("udp", s.Listen + ":53")
	if err != nil {
		return err
	}
	defer pc.Close()


	go func() {
		for {
			buffer := make([]byte, bufferSize)
			n, addr, err := pc.ReadFrom(buffer)
			if err != nil {
				return
			}
			go s.handlePacket(pc, addr, buffer[:n])
		}
	}()

	select {
	case <-s.Ctx.Done():
		err = s.Ctx.Err()
	}
	return nil
}

func (s Server) handlePacket(pc net.PacketConn, remote net.Addr, data []byte) {
	query := dns.Msg{}
	err := query.Unpack(data)
	if err != nil {
		return
	}
	if len(query.Question) == 1 {
		q := query.Question[0]
		log.Printf("[*] [%s] Query [%s] [%s]\n", remote.String(), dns.Type(q.Qtype).String(), q.Name)
	}
	result, err := s.Client.Query(data)
	if err != nil {
		log.Printf("[!] ERROR - %s\n", err.Error())
		return
	}
	pc.WriteTo(result, remote)

}
