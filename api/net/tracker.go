package net

import (
	"fmt"

	"github.com/science-engineering-art/gotify/tracker/core"
)

var (
	Tracker *core.EmptyTracker
)

func InitTracker(ip string, port int, bootPort int, isBoot bool) {
	Tracker, _ = core.NewEmptyTracker(ip, port, bootPort, false)
	addr := fmt.Sprintf("%s:%d", ip, port)
	go Tracker.FullNode.CreateGRPCServer(addr)
}
