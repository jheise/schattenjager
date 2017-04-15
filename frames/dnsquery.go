package frames

import (
	//standard lib
	"net"
)

type DNSQuery struct {
	SrcIP, DstIP net.IP
	Query        string
	Type         string
	Request      bool
	Timestamp    int64
	Domain       string
	Subdomain    string
	format       string
}
