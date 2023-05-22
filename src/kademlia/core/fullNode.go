package core

import (
	"crypto/sha1"
	"fmt"

	"github.com/science-engineering-art/spotify/src/kademlia/interfaces"
	"github.com/science-engineering-art/spotify/src/kademlia/structs"
)

type FullNode struct {
	structs.Node
	RoutingTable *structs.RoutingTable
	Storage      interfaces.Persistence
}

// TODO: NewFullNode method for node initializing

// TODO: JoinNetwork method for connecting to a bootstrap node

func (fn *FullNode) Ping(sender structs.Node) structs.Node {
	err := fn.RoutingTable.AddNode(sender)
	if err != nil {
		fmt.Println("An error ocurred when executing Ping: ", err)
	}
	return fn.Node
}

func (fn *FullNode) Store(data *[]byte) error {
	sha := sha1.Sum(*data)

	err := fn.Storage.Create(sha[:], data)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Check Find Value return when the value is not in the node
func (fn *FullNode) FindValue(infoHash *[]byte) (value *[]byte) {
	value, err := fn.Storage.Read(*infoHash)
	if err != nil {
		fmt.Println("Find Value error: ", err)
	}
	return value
}

// TODO: Check get k nearest values return
// func (fn *FullNode) FindNode(target *[]byte) (kBucket *[]structs.Node) {
// 	if bytes.Compare(fn.ID, *target) == 0 {
// 		kBucket = &[]structs.Node{fn.Node}
// 	}
// 	sl := fn.RoutingTable.GetClosestContacts(3, *target, []*structs.Node{&fn.Node})

// 	//kBucket = sl.Nodes
// 	return
// }
