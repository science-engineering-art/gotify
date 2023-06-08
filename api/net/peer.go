package net

import (
	"os"

	"github.com/science-engineering-art/gotify/peer/core"
)

var (
	Peer *core.Peer
)

func InitPeer() {
	Peer = core.NewPeer(os.Getenv("MONGODB_IP"), false)
	go Peer.FullNode.CreateGRPCServer("0.0.0.0:8080")
}
