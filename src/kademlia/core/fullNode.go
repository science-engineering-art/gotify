package core

import (
	"context"
	"crypto/rand"

	"github.com/science-engineering-art/spotify/src/kademlia/interfaces"
	"github.com/science-engineering-art/spotify/src/kademlia/pb"
	"github.com/science-engineering-art/spotify/src/kademlia/structs"
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

func NewFullNode(ip string, port int, storage interfaces.Persistence) *FullNode {
	id, _ := newID()
	node := structs.Node{ID: id, IP: ip, Port: port}
	routingTable := structs.RoutingTable{}
	dht := DHT{Node: node, RoutingTable: &routingTable, Storage: storage}
	fullNode := FullNode{dht: dht}
	return &fullNode
}

// newID generates a new random ID
func newID() ([]byte, error) {
	result := make([]byte, 20)
	_, err := rand.Read(result)
	return result, err
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

func (fn *FullNode) LookUp(action int, target []byte, data *[]byte) error {

	// sl := fn.dht.RoutingTable.GetClosestContacts(structs.Alpha, target, []*structs.Node{})

	// contacted := make(map[string]bool)

	// // We keep track of nodes contacted so far. We don't contact the same node
	// // twice.
	// var contacted = make(map[string]bool)

	// // According to the Kademlia white paper, after a round of FIND_NODE RPCs
	// // fails to provide a node closer than closestNode, we should send a
	// // FIND_NODE RPC to all remaining nodes in the shortlist that have not
	// // yet been contacted.
	// queryRest := false

	// // We keep a reference to the closestNode. If after performing a search
	// // we do not find a closer node, we stop searching.
	// if len(sl.Nodes) == 0 {
	// 	return nil, nil, nil
	// }

	// closestNode := sl.Nodes[0]

	// if t == iterateFindNode {
	// 	bucket := getBucketIndexFromDifferingBit(target, dht.ht.Self.ID)
	// 	dht.ht.resetRefreshTimeForBucket(bucket)
	// }

	// removeFromShortlist := []*NetworkNode{}

	// for {
	// 	expectedResponses := []*expectedResponse{}
	// 	numExpectedResponses := 0

	// 	// Next we send messages to the first (closest) alpha nodes in the
	// 	// shortlist and wait for a response

	// 	for i, node := range sl.Nodes {
	// 		// Contact only alpha nodes
	// 		if i >= alpha && !queryRest {
	// 			break
	// 		}

	// 		// Don't contact nodes already contacted
	// 		if contacted[string(node.ID)] {
	// 			continue
	// 		}

	// 		contacted[string(node.ID)] = true
	// 		query := &message{}
	// 		query.Sender = dht.ht.Self
	// 		query.Receiver = node

	// 		switch t {
	// 		case iterateFindNode:
	// 			query.Type = messageTypeFindNode
	// 			queryData := &queryDataFindNode{}
	// 			queryData.Target = target
	// 			query.Data = queryData
	// 		case iterateFindValue:
	// 			query.Type = messageTypeFindValue
	// 			queryData := &queryDataFindValue{}
	// 			queryData.Target = target
	// 			query.Data = queryData
	// 		case iterateStore:
	// 			query.Type = messageTypeFindNode
	// 			queryData := &queryDataFindNode{}
	// 			queryData.Target = target
	// 			query.Data = queryData
	// 		default:
	// 			panic("Unknown iterate type")
	// 		}

	// 		// Send the async queries and wait for a response
	// 		res, err := dht.networking.sendMessage(query, true, -1)
	// 		if err != nil {
	// 			// Node was unreachable for some reason. We will have to remove
	// 			// it from the shortlist, but we will keep it in our routing
	// 			// table in hopes that it might come back online in the future.
	// 			removeFromShortlist = append(removeFromShortlist, query.Receiver)
	// 			continue
	// 		}

	// 		expectedResponses = append(expectedResponses, res)
	// 	}

	// 	for _, n := range removeFromShortlist {
	// 		sl.RemoveNode(n)
	// 	}

	// 	numExpectedResponses = len(expectedResponses)

	// 	resultChan := make(chan (*message))
	// 	for _, r := range expectedResponses {
	// 		go func(r *expectedResponse) {
	// 			select {
	// 			case result := <-r.ch:
	// 				if result == nil {
	// 					// Channel was closed
	// 					return
	// 				}
	// 				dht.addNode(newNode(result.Sender))
	// 				resultChan <- result
	// 				return
	// 			case <-time.After(dht.options.TMsgTimeout):
	// 				dht.networking.cancelResponse(r)
	// 				return
	// 			}
	// 		}(r)
	// 	}

	// 	var results []*message
	// 	if numExpectedResponses > 0 {
	// 	Loop:
	// 		for {
	// 			select {
	// 			case result := <-resultChan:
	// 				if result != nil {
	// 					results = append(results, result)
	// 				} else {
	// 					numExpectedResponses--
	// 				}
	// 				if len(results) == numExpectedResponses {
	// 					close(resultChan)
	// 					break Loop
	// 				}
	// 			case <-time.After(dht.options.TMsgTimeout):
	// 				close(resultChan)
	// 				break Loop
	// 			}
	// 		}

	// 		// TODO
	// 		// leandro_driguez: en algún momento se lleva a este código?
	// 		for _, result := range results {
	// 			if result.Error != nil {
	// 				sl.RemoveNode(result.Receiver)
	// 				continue
	// 			}
	// 			switch t {
	// 			case iterateFindNode:
	// 				responseData := result.Data.(*responseDataFindNode)
	// 				sl.AppendUniqueNetworkNodes(responseData.Closest)
	// 			case iterateFindValue:
	// 				responseData := result.Data.(*responseDataFindValue)
	// 				// TODO When an iterativeFindValue succeeds, the initiator must
	// 				// store the key/value pair at the closest node seen which did
	// 				// not return the value.
	// 				if responseData.Value != nil {
	// 					return responseData.Value, nil, nil
	// 				}
	// 				sl.AppendUniqueNetworkNodes(responseData.Closest)
	// 			case iterateStore:
	// 				responseData := result.Data.(*responseDataFindNode)
	// 				sl.AppendUniqueNetworkNodes(responseData.Closest)
	// 			}
	// 		}
	// 	}

	// 	if !queryRest && len(sl.Nodes) == 0 {
	// 		return nil, nil, nil
	// 	}

	// 	sort.Sort(sl)

	// 	// If closestNode is unchanged then we are done
	// 	if bytes.Equal(sl.Nodes[0].ID, closestNode.ID) || queryRest {
	// 		// We are done
	// 		switch t {
	// 		case iterateFindNode:
	// 			if !queryRest {
	// 				queryRest = true
	// 				continue
	// 			}
	// 			return nil, sl.Nodes, nil
	// 		case iterateFindValue:
	// 			return nil, sl.Nodes, nil
	// 		case iterateStore:
	// 			for i, n := range sl.Nodes {
	// 				if i >= k {
	// 					return nil, nil, nil
	// 				}

	// 				query := &message{}
	// 				query.Receiver = n
	// 				query.Sender = dht.ht.Self
	// 				query.Type = messageTypeStore
	// 				queryData := &queryDataStore{}
	// 				queryData.Data = data
	// 				query.Data = queryData
	// 				dht.networking.sendMessage(query, false, -1)
	// 			}
	// 			return nil, nil, nil
	// 		}
	// 	} else {
	// 		closestNode = sl.Nodes[0]
	// 	}
	// }

	return nil
}
