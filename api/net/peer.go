package net

import (
	"github.com/science-engineering-art/gotify/peer/core"
)

var (
	Peer *core.EmptyPeer
)

func InitPeer() {
	Peer = core.NewEmptyPeer(false)
	go Peer.FullNode.CreateGRPCServer("0.0.0.0:8080")
}
