package util

import (
	"context"
	"fmt"
	"net"

	"github.com/mat-penna/GoDNSFilter/data"
)

func GoogleDNSDialer(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{}
	return d.DialContext(ctx, "udp", "8.8.8.8:53")
}

func CloudflareDNSDialer(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{}
	return d.DialContext(ctx, "udp", "1.1.1.1:53")
}

func QueryIPfromURL(url string) string {

	r := net.Resolver{
		PreferGo: true,
		Dial:     CloudflareDNSDialer,
	}
	ctx := context.Background()
	ipaddr, err := r.LookupIPAddr(ctx, url)
	if err != nil {
		return data.SinkholeAddr
	}
	fmt.Println("DNS Result", ipaddr[0].IP.String())
	if len(ipaddr) > 0 {
		return ipaddr[0].IP.String()
	} else {
		return data.SinkholeAddr
	}
}
