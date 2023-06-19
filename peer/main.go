package main

import (
	"fmt"

	"github.com/science-engineering-art/gotify/peer/core"
	"github.com/science-engineering-art/gotify/peer/utils"
)

var (
	ip   = utils.GetIpFromHost()
	port = 8080
)

func main() {
	peer := core.NewRedisPeer(ip, port, 32140, true)

	grpcAddr := fmt.Sprintf("%s:%d", ip, port)

	peer.CreateGRPCServer(grpcAddr)
}
