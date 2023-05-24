package core

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"

	"github.com/science-engineering-art/spotify/src/kademlia/interfaces"
	"github.com/science-engineering-art/spotify/src/kademlia/pb"
	"github.com/science-engineering-art/spotify/src/kademlia/structs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	StoreValue = iota
	KNeartestNodes
	GetValue
)

type FullNode struct {
	pb.UnimplementedFullNodeServer
	dht DHT
}

func NewGrpcFullNodeServer(ip string, port int, storage interfaces.Persistence) *FullNode {
	id, _ := NewID(ip, port)
	node := structs.Node{ID: id, IP: ip, Port: port}
	routingTable := structs.NewRoutingTable(node)
	dht := DHT{Node: node, RoutingTable: routingTable, Storage: storage}
	fullNode := FullNode{dht: dht}
	return &fullNode
}

// newID generates a new random ID
func NewID(ip string, port int) ([]byte, error) {
	hashValue := sha1.Sum([]byte(ip + ":" + strconv.FormatInt(int64(port), 10)))
	return []byte(hashValue[:]), nil
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

func (fn *FullNode) FindNode(ctx context.Context, target *pb.TargetID) (*pb.KBucket, error) {
	bucket := fn.dht.FindNode(&target.ID)
	return getKBucketFromNodeArray(bucket), nil
}

func (fn *FullNode) FindValue(target *pb.TargetID, fv pb.FullNode_FindValueServer) error {
	value, neighbors := fn.dht.FindValue(&target.ID)
	kbucket := getKBucketFromNodeArray(neighbors)
	response := pb.FindValueResponse{KNeartestBuckets: kbucket, Value: &pb.Data{Init: 0, End: int32(4000024), Buffer: (*value)[:4000024]}}
	fv.Send(&response)
	return nil
}

func getKBucketFromNodeArray(nodes *[]structs.Node) *pb.KBucket {
	result := &pb.KBucket{}
	for _, node := range *nodes {
		(*result).Bucket = append((*result).Bucket, &pb.Node{ID: node.ID, IP: node.IP, Port: int32(node.Port)})
	}
	return result
}

func (fn *FullNode) LookUp(action int, target []byte, data *[]byte) (*[]byte, []structs.Node, error) {

	sl := fn.dht.RoutingTable.GetClosestContacts(structs.Alpha, target, []*structs.Node{})

	contacted := make(map[string]bool)

	if len(*sl.Nodes) == 0 {
		return nil, nil, nil
	}

	for {
		addedNodes := 0

		for i, node := range *sl.Nodes {
			if i >= structs.Alpha {
				break
			}
			if contacted[string(node.ID)] {
				continue
			}
			contacted[string(node.ID)] = true

			// get RPC client
			address := fmt.Sprintf("%s:%d", node.IP, node.Port)
			conn, _ := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
			client := pb.NewFullNodeClient(conn)

			// function to add the received nodes into the short list
			addRecvNodes := func(recvNodes *pb.KBucket) {
				kBucket := []*structs.Node{}
				for _, pbNode := range recvNodes.Bucket {
					if !contacted[string(pbNode.ID)] {
						kBucket = append(kBucket, &structs.Node{
							ID:   pbNode.ID,
							IP:   pbNode.IP,
							Port: int(pbNode.Port),
						})
						addedNodes++
					}
				}
				sl.Append(kBucket)
			}

			switch action {
			case StoreValue:
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

				recvNodes, err := client.FindNode(ctx, &pb.TargetID{ID: node.ID})
				if err != nil {
					return nil, nil, err
				}
				addRecvNodes(recvNodes)
			case KNeartestNodes:
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

				recvNodes, err := client.FindNode(ctx, &pb.TargetID{ID: node.ID})
				if err != nil {
					return nil, nil, err
				}
				addRecvNodes(recvNodes)
			case GetValue:
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

				stream, err := client.FindValue(ctx, &pb.TargetID{ID: node.ID})
				if err != nil {
					return nil, nil, err
				}

				buffer := []byte{}
				var init int32 = 0

				for {
					resp, err := stream.Recv()
					if resp == nil {
						break
					}
					if err != nil {
						return nil, nil, err
					}

					if resp.KNeartestBuckets != nil && resp.Value == nil {
						addRecvNodes(resp.KNeartestBuckets)
						continue
					}

					if resp.KNeartestBuckets == nil && resp.Value != nil {
						if init == resp.Value.Init {
							buffer = append(buffer, resp.Value.Buffer...)
							init = resp.Value.End
						} else {
							return nil, nil, errors.New("error")
						}
					}
				}

				if len(buffer) > 0 {
					return &buffer, nil, nil
				}
			}

		}

		sl.Comparator = fn.dht.ID
		sort.Sort(sl)

		if addedNodes == 0 {
			break
		}
	}

	kBucket := []structs.Node{}

	for i, node := range *sl.Nodes {
		if i == structs.K {
			break
		}
		switch action {
		case StoreValue:
			address := fmt.Sprintf("%s:%d", node.IP, node.Port)
			conn, _ := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
			client := pb.NewFullNodeClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			stream, err := client.Store(ctx)
			if err != nil {
				return nil, nil, err
			}

			for i := 0; i < len(*data); i++ {
				init := int32(i)
				end := int32(math.Max(float64(i+1024), float64(len(*data))))

				stream.Send(&pb.Data{
					Init:   init,
					End:    end,
					Buffer: (*data)[init:end],
				})
			}
			return nil, nil, nil
		case KNeartestNodes:
			kBucket = append(kBucket, structs.Node{
				ID:   node.ID,
				IP:   node.IP,
				Port: node.Port,
			})
		}
	}
	return nil, kBucket, nil
}
