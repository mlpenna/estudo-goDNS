package dns_logger

import (
	"fmt"
	"net"
)

type DNSQueryEntry struct {
	TimeDate     string
	SrcHost      net.Addr
	UrlQuery     string
	ReturnedAdrr string
}

func PrintQuery(d DNSQueryEntry) {
	fmt.Println(d.SrcHost.String() + "," + d.TimeDate + "," + d.UrlQuery + "," + d.ReturnedAdrr)
}
