package net

import (
	"os"

	"github.com/science-engineering-art/gotify/peer/core"
	kademlia "github.com/science-engineering-art/kademlia-grpc/core"
)

var (
	Peer *core.Peer
)

func InitPeer() {
	Peer = core.NewPeer(os.Getenv("MONGODB_IP"), false)
	go kademlia.CreateGRPCServerFromFullNode(Peer.FullNode, "0.0.0.0:8080")
}
