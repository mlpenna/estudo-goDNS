package logger

import (
	"fmt"
	"net"
)

type DNSQueryEntry struct {
	TimeDate     string
	SrcHost      net.Addr
	UrlQuery     string
	ReturnedAdrr string
	Domain       string
}

func PrintQuery(d DNSQueryEntry) {
	fmt.Println(d.SrcHost.String() + "," + d.TimeDate + "," + d.UrlQuery + "," + d.ReturnedAdrr + "," + d.Domain)
}
