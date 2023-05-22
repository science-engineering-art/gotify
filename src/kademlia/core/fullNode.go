package core

import (
	"fmt"

	"github.com/science-engineering-art/spotify/src/kademlia/structs"
)

type FullNode struct {
	structs.Node
	RoutingTable *structs.RoutingTable
	// StorageTable Storage
}

// TODO: NewFullNode method for node initializing

// TODO: JoinNetwork method for connecting to a bootstrap node

func (n *FullNode) Ping(sender structs.Node) structs.Node {
	err := n.RoutingTable.AddNode(sender)
	if err != nil {
		fmt.Println("An error ocurred when executing Ping: ", err)
	}
	return n.Node
}
