package core

import "github.com/science-engineering-art/spotify/src/kademlia/structs"

type FullNode struct {
	structs.Node
	RoutingTable map[string][]structs.Node
	// StorageTable Storage
}
