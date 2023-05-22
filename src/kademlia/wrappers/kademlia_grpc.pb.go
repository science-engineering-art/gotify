package wrappers

import (
	"context"

	"github.com/science-engineering-art/spotify/src/kademlia/pb"
)

type KademliaProtocol struct {
	pb.UnimplementedKademliaProtocolServer
}

func (kp *KademliaProtocol) Ping(context.Context, *pb.Node) (*pb.Node, error) {

	return nil, nil
}

func (kp *KademliaProtocol) Store(pb.KademliaProtocol_StoreServer) error {
	return nil
}

func (kp *KademliaProtocol) FindNode(context.Context, *pb.TargetID) (*pb.KBucket, error) {
	return nil, nil
}

func (kp *KademliaProtocol) FindValue(*pb.TargetID, pb.KademliaProtocol_FindValueServer) error {
	return nil
}
