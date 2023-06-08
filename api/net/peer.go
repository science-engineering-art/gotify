package net

import (
	"fmt"

	"github.com/science-engineering-art/gotify/peer/core"
)

var (
	Peer *core.EmptyPeer
)

func InitPeer(ip string, port int) {
	Peer = core.NewEmptyPeer(ip, false)
	addr := fmt.Sprintf("%s:%d", ip, port)
	go Peer.FullNode.CreateGRPCServer(addr)
}
