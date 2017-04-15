package frames

import (
	//standard lib
	"net"
)

type DNSCapture struct {
	SrcIP, DstIP net.IP
	Answers      []net.IP
	Query        string
	Request      bool
	Timestamp    int64
}

type DNSFrame interface {
	Format() string
	SrcIP() net.IP
	DstIP() net.IP
	Timestamp() int64
}
