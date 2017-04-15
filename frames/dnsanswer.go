package frames

import (
	//standard lib
	"net"
)

type DNSAnswer struct {
	SrcIP, DstIP net.IP
	Answer       string
	City         string
	Country      string
	Domain       string
	Geoip        string
	Hash         string
	Subdomain    string
	Type         string
	Query        string
	Request      bool
	Timestamp    int64
	TTL          uint32
	ASN          uint32
}
