package net

import (
	"github.com/science-engineering-art/gotify/tracker/core"
)

var (
	Tracker *core.Tracker
)

func InitTracker(ip string, port int, bootPort int, isBoot bool) {
	Tracker, _ = core.NewTracker(ip, port, bootPort, isBoot)
	go Tracker.FullNode.CreateGRPCServer("0.0.0.0:8080")
}
