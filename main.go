package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/miekg/dns"
)

var domainsToBlock map[string]string = map[string]string{
	"google.com.":   "1.2.3.4",
	"facebook.com.": "1.2.3.4",
}

type handler struct{}

func (*handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)
	QueryIP(msg.Question[0].Name)
	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		domain := msg.Question[0].Name
		address, ok := domainsToBlock[domain]
		if ok {
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
				A:   net.ParseIP(QueryIP(address)),
			})
		} else {
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
				A:   net.ParseIP(QueryIP(msg.Question[0].Name)),
			})
		}
	}
	// fmt.Println(msg)
	w.WriteMsg(&msg)
}

func QueryIP(url string) string {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, "1.1.1.1:53")
		},
	}
	ip, _ := r.LookupHost(context.Background(), url)

	// print(ip[0])
	if len(ip) == 0 {
		return "1.2.3.4"
	} else {
		return ip[0]
	}

}

func main() {
	srv := &dns.Server{Addr: ":" + strconv.Itoa(533), Net: "udp"}
	srv.Handler = &handler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}

}
