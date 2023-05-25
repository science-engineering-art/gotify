package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp4", "255.255.255.255:3999")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("upd4", addr)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 1024)
	n, rAddr, err := conn.ReadFrom(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Received %d bytes from %v", n, rAddr)
}
