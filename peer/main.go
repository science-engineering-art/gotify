package main

import (
	kademlia "github.com/science-engineering-art/spotify/src/kademlia/core"
	"github.com/science-engineering-art/spotify/src/peer/persistence"
)

var (
	db = persistence.NewMongoDb("admin", "songs")
)

func main() {
	_ = kademlia.NewFullNode("localhost", 8080, db)
}
