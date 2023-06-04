package main

import (
	"github.com/science-engineering-art/gotify/peer/persistence"
	kademlia "github.com/science-engineering-art/kademlia-grpc/core"
)

var (
	db = persistence.NewMongoDb("admin", "songs")
)

func main() {
	_ = kademlia.NewFullNode("0.0.0.0", 8080, 32140, db, false)
}
