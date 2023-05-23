package wrappers

import (
	"context"

	"github.com/science-engineering-art/spotify/src/kademlia/core"
	"github.com/science-engineering-art/spotify/src/kademlia/pb"
)

type FullNode struct {
	pb.UnimplementedFullNodeServer
	DHT core.DHT
}

func (kp *FullNode) Ping(context.Context, *pb.Node) (*pb.Node, error) {

	return nil, nil
}

func (kp *FullNode) Store(pb.FullNode_StoreServer) error {
	return nil
}

func (kp *FullNode) FindNode(context.Context, *pb.TargetID) (*pb.KBucket, error) {
	return nil, nil
}

func (kp *FullNode) FindValue(*pb.TargetID, pb.FullNode_FindValueServer) error {
	return nil
}
