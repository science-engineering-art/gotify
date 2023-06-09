package net

import (
	"fmt"

	"github.com/science-engineering-art/gotify/tracker/core"
)

var (
	Tracker *core.Tracker
)

func InitTracker(ip string, port int, bootPort int, isBoot bool) {
	Tracker, _ = core.NewTracker(ip, port, bootPort, isBoot)
	addr := fmt.Sprintf("%s:%d", ip, port)
	go Tracker.FullNode.CreateGRPCServer(addr)
}
