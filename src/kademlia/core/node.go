package core

import s "github.com/science-engineering-art/spotify/src/kademlia/structs"

type FullNode struct {
	s.Bucket
	RoutingTable map[string][]s.Bucket
	// StorageTable Storage
}
