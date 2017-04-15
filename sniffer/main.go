package main

import (

	//external
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	/*
		addr    = kingpin.Arg("addr", "address to bind to").Default("*").String()
		port    = kingpin.Arg("port", "port to expose for zmq").Default("7777").String()
	*/

	iface     = kingpin.Arg("interface", "interface to sniff traffic on").Default("wlp1s0").String()
	configsrc = kingpin.Arg("config", "Config file to load").Default("sniffer.conf").String()
	verbose   bool
)

func main() {
	kingpin.Parse()
	caps := make(chan interface{})
	packets := make(chan gopacket.Packet)

	// process config
	config, err := ReadConfig(*configsrc)
	if err != nil {
		panic(err)
	}

	verbose = config.Verbose

	// start packet processor
	go processer(packets, caps)

	var publisher Publisher

	// start capture consumer
	if config.Publisher == "zmq" {
		publisher, err = newZMQPublisher(config.PubConfig["address"].(string),
			config.PubConfig["port"].(string),
			caps)
		if err != nil {
			panic(err)
		}
	}

	go publisher.Process()

	handle, err := pcap.OpenLive(*iface, 1600, true, 0)
	if err != nil {
		panic(err)
	}
	err = handle.SetBPFFilter("udp and port 53")
	if err != nil {
		panic(err)
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		packets <- packet
	}
}
