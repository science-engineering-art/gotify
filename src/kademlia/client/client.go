package main

import (
	"log"
	"net"
)

func main() {
	addr := net.UDPAddr{
		IP:   net.IPv4(255, 255, 255, 255),
		Port: 8888,
	}

	conn, err := net.DialUDP("udp4", nil, &addr)
	if err != nil {
		log.Fatal(err)
	}

	buffer := []byte("Hello world!")
	_, err = conn.Write(buffer)
	if err != nil {
		log.Fatal(err)
	}
}
