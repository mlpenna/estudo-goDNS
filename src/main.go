package main

//dig @127.0.0.1 -p 533 google.com

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/mat-penna/GoDNSFilter/data"
	"github.com/mat-penna/GoDNSFilter/logger"
	"github.com/mat-penna/GoDNSFilter/util"

	"github.com/miekg/dns"
)

type handler struct{}

var IPad string
var domain string

func (*handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)
	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		domain = msg.Question[0].Name
		block := data.IsInBlockList(domain)
		if block {
			IPad = data.SinkholeAddr
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
				A:   net.ParseIP(IPad),
			})
		} else {
			IPad = util.QueryIPfromURL(domain)
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
				A:   net.ParseIP(IPad),
			})
		}
	}
	w.WriteMsg(&msg)

	var e logger.DNSQueryEntry
	e.SrcHost = w.RemoteAddr()
	e.TimeDate = time.Now().String()
	e.UrlQuery = IPad
	e.Domain = domain

	logger.PrintQuery(e)
}

func main() {

	// data.AddDomainToBlock("google.com.")
	data.AddDomainToBlock("facebook.com.")

	fmt.Println("Iniciando Servidor DNS...")
	srv := &dns.Server{Addr: ":" + data.LocalPort, Net: "udp"}
	srv.Handler = &handler{}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}
