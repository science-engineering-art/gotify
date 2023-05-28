package core

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"net"
	"sort"
	"strconv"
	"time"

	"github.com/science-engineering-art/spotify/src/kademlia/interfaces"
	"github.com/science-engineering-art/spotify/src/kademlia/pb"
	"github.com/science-engineering-art/spotify/src/kademlia/structs"
	"github.com/science-engineering-art/spotify/src/kademlia/utils"
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
	dht *DHT
}

func NewFullNode(ip string, port, bootstrapPort int, storage interfaces.Persistence, isBootstrapNode bool) *FullNode {

	id, _ := NewID(ip, port)
	node := structs.Node{ID: id, IP: ip, Port: port}
	routingTable := structs.NewRoutingTable(node)
	dht := DHT{Node: node, RoutingTable: routingTable, Storage: storage}
	fullNode := FullNode{dht: &dht}

	if isBootstrapNode {
		go listen(bootstrapPort, &dht)
	}

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
	//fmt.Printf("The value that reached %s", b58.Encode(buffer))
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
	//fmt.Println("Retrieved value:", b58.Encode(*value))
	kbucket := &pb.KBucket{Bucket: []*pb.Node{}}
	if neighbors != nil && len(*neighbors) > 0 {
		kbucket = getKBucketFromNodeArray(neighbors)
	}
	if value == nil {
		value = &[]byte{}
	}
	response := pb.FindValueResponse{KNeartestBuckets: kbucket, Value: &pb.Data{Init: 0, End: int32(len(*value)), Buffer: (*value)[:]}}
	fv.Send(&response)
	return nil
}

func getKBucketFromNodeArray(nodes *[]structs.Node) *pb.KBucket {
	result := pb.KBucket{Bucket: []*pb.Node{}}
	for _, node := range *nodes {
		//fmt.Println("In the for:", node)
		result.Bucket = append(result.Bucket, &pb.Node{ID: node.ID, IP: node.IP, Port: int32(node.Port)})
		//fmt.Println("Append Well")
	}
	return &result
}

func (fn *FullNode) LookUp(target []byte) ([]structs.Node, error) {

	sl := fn.dht.RoutingTable.GetClosestContacts(structs.Alpha, target, []*structs.Node{})

	contacted := make(map[string]bool)

	if len(*sl.Nodes) == 0 {
		return nil, nil
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

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			recvNodes, err := client.FindNode(ctx, &pb.TargetID{ID: node.ID})
			if err != nil {
				return nil, err
			}
			addRecvNodes(recvNodes)
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
		kBucket = append(kBucket, structs.Node{
			ID:   node.ID,
			IP:   node.IP,
			Port: node.Port,
		})
	}
	return kBucket, nil
}

func listen(port int, dht *DHT) {
	strAddr := fmt.Sprintf("0.0.0.0:%d", port)

	addr, err := net.ResolveUDPAddr("udp4", strAddr)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 1024)

	for {
		n, rAddr, err := conn.ReadFrom(buffer)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Received %d bytes from %v\n", n, rAddr)

		go func() {
			kBucket := dht.FindNode(&buffer)

			host, port, _ := net.SplitHostPort(rAddr.String())
			portInt, _ := strconv.Atoi(port)

			dht.RoutingTable.AddNode(structs.Node{IP: host, Port: portInt})

			respConn, err := net.Dial("tcp", rAddr.String())
			if err != nil {
				panic(err)
			}
			defer conn.Close()

			bytesKBucket, err := utils.SerializeMessage(kBucket)
			if err != nil {
				panic(err)
			}

			respConn.Write(*bytesKBucket)
		}()
	}
}
