package core

import s "github.com/science-engineering-art/spotify/src/kademlia/structs"

type FullNode struct {
	s.Node
	RoutingTable map[string][]s.Node
	// StorageTable Storage
}
