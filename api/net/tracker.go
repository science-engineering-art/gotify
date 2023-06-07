package net

import (
	"github.com/science-engineering-art/gotify/tracker/core"
	kademlia "github.com/science-engineering-art/kademlia-grpc/core"
)

var (
	Tracker *core.Tracker
)

func InitTracker(ip string, port int, bootPort int, isBoot bool) {
	Tracker, _ = core.NewTracker(ip, port, bootPort, isBoot)
	go kademlia.CreateGRPCServerFromFullNode(Tracker.FullNode, "0.0.0.0:8080")
}
