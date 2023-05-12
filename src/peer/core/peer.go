package core

import s "github.com/science-engineering-art/spotify/src/peer/core/structs"

type Peer struct {
	s.Node
	RoutingTable map[string][]s.Bucket
	StorageTable map[string]string
}
