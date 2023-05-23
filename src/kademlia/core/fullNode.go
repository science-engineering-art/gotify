package core

import (
	"context"

	"github.com/science-engineering-art/spotify/src/kademlia/pb"
	"github.com/science-engineering-art/spotify/src/kademlia/structs"
)

type FullNode struct {
	pb.UnimplementedFullNodeServer
	dht DHT
}

func (fn *FullNode) Ping(ctx context.Context, sender *pb.Node) (*pb.Node, error) {

	err := fn.dht.RoutingTable.AddNode(
		structs.Node{
			ID:   sender.ID,
			IP:   sender.IP,
			Port: int(sender.Port),
		})
	if err != nil {
		return nil, err
	}

	receiver := &pb.Node{ID: fn.dht.ID, IP: fn.dht.IP, Port: int32(fn.dht.Port)}
	return receiver, nil
}

func (fn *FullNode) Store(stream pb.FullNode_StoreServer) error {
	buffer := []byte{}
	var init int32 = 0

	for {
		data, err := stream.Recv()
		if data == nil {
			break
		}

		if init == data.Init {
			buffer = append(buffer, data.Buffer...)
			init = data.End
		} else {
			return err
		}

		if err != nil {
			return err
		}
	}

	err := fn.dht.Store(&buffer)
	if err != nil {
		return err
	}
	return nil
}

func (fn *FullNode) FindNode(context.Context, *pb.TargetID) (*pb.KBucket, error) {
	return nil, nil
}

func (fn *FullNode) FindValue(*pb.TargetID, pb.FullNode_FindValueServer) error {
	return nil
}
