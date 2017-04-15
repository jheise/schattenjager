package main

import (
	//standard lib
	"encoding/json"
	"fmt"

	//external
	zmq "github.com/pebbe/zmq4"
)

type ZMQPublisher struct {
	socket  *zmq.Socket
	ipaddr  string
	port    string
	connect string
	caps    chan interface{}
}

func newZMQPublisher(ipaddr string, port string, caps chan interface{}) (*ZMQPublisher, error) {
	newpub := new(ZMQPublisher)
	socket, err := zmq.NewSocket(zmq.PUB)
	if err != nil {
		return newpub, err
	}

	newpub.socket = socket
	newpub.ipaddr = ipaddr
	newpub.port = port
	newpub.connect = fmt.Sprintf("tcp://%s:%s", ipaddr, port)
	newpub.caps = caps

	newpub.socket.Bind(newpub.connect)

	return newpub, nil
}

func (self *ZMQPublisher) Process() {

	for cap := range self.caps {
		cap_json, err := json.Marshal(cap)
		if err != nil {
			panic(err)
		}
		if verbose {
			fmt.Printf("%s\n", cap_json)
		}
		msg := fmt.Sprintf("dns %s", cap_json)
		self.socket.SendMessage(msg)
	}
}
