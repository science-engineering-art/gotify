package main

import (
	"fmt"

	trackerCore "github.com/science-engineering-art/gotify/tracker/core"
	"github.com/science-engineering-art/gotify/tracker/utils"
)

var (
	ip   = utils.GetIpFromHost()
	port = 9090
)

func main() {
	tracker, _ := trackerCore.NewTracker(ip, port, 42140, true)

	grpcAddr := fmt.Sprintf("%s:%d", ip, port)

	tracker.CreateGRPCServer(grpcAddr)
}
