package structs

import (
	"bytes"
	"context"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/science-engineering-art/spotify/src/kademlia/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// a small number representing the degree of parallelism in network calls
	alpha = 3

	// the size in bits of the keys used to identify nodes and store and
	// retrieve data; in basic Kademlia this is 160, the length of a SHA1
	b = 160

	// the maximum number of contacts stored in a bucket
	k = 20
)

type RoutingTable struct {
	NodeInfo Node
	KBuckets [][]Node
	mutex    *sync.Mutex
}

func (rt *RoutingTable) init(b Node) {
	rt.NodeInfo = b
	rt.KBuckets = [][]Node{}
	rt.mutex = &sync.Mutex{}
}

func (rt *RoutingTable) isAlive(b Node) bool {
	address := fmt.Sprintf("%s:%d", rt.NodeInfo.IP, rt.NodeInfo.Port)
	conn, _ := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	client := pb.NewKademliaProtocolClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pbNode, err := client.Ping(ctx, &pb.Node{})
	if err != nil {
		return false
	}

	if !rt.NodeInfo.Equal(Node{ID: pbNode.ID, IP: pbNode.IP, Port: int(pbNode.Port)}) {
		return false
	}

	return true
}

// Función que se encarga de añadir un nodo a la tabla de
// rutas con las restricciones pertinentes del protocolo
func (rt *RoutingTable) AddNode(b Node) error {
	rt.mutex.Lock()
	defer rt.mutex.Unlock()

	// get the correspondient bucket
	bIndex := getBucketIndex(rt.NodeInfo.ID, b.ID)
	bucket := rt.KBuckets[bIndex]

	// update the node position in the case of it already
	// belongs to the bucket
	for i := 0; i < len(bucket); i++ {
		if bucket[i].Equal(b) {
			bucket = append(bucket[:i], bucket[i+1:]...)
			bucket = append(bucket, b)
			goto RETURN
		}
	}

	if len(bucket) < k {
		bucket = append(bucket, b)
	} else if !rt.isAlive(bucket[0]) {
		bucket = append(bucket[1:], b)
	}

RETURN:
	rt.KBuckets[bIndex] = bucket
	return nil
}

// TODO
// leandro_driguez: chequear bien el uso de este método,
// no vaya a ser que se esté utilizando mal
func getBucketIndex(id1 []byte, id2 []byte) int {

	// Look at each byte from left to right
	for j := 0; j < len(id1); j++ {
		// xor the byte
		xor := id1[j] ^ id2[j]

		// check each bit on the xored result from left to right in order
		for i := 0; i < 8; i++ {
			if hasBit(xor, uint(i)) {
				byteIndex := j * 8
				bitIndex := i
				// return b - (byteIndex + bitIndex) - 1 // leandro_driguez: no será (byteIndex + bitIndex) solamente?
				return byteIndex + bitIndex
			}
		}
	}

	// the ids must be the same
	// this should only happen during bootstrapping
	return 0
}

// Simple helper function to determine the value of a particular
// bit in a byte by index
//
// Example:
// number:  1
// bits:    00000001
// pos:     01234567
func hasBit(n byte, pos uint) bool {
	pos = 7 - pos
	val := n & (1 << pos)
	return (val > 0)
}

func (rt *RoutingTable) getClosestContacts(num int, target []byte, ignoredNodes []*Node) *ShortList {
	rt.mutex.Lock()
	defer rt.mutex.Unlock()
	// First we need to build the list of adjacent indices to our target
	// in order
	index := getBucketIndex(rt.NodeInfo.ID, target)
	indexList := []int{index}
	i := index - 1
	j := index + 1

	// TODO
	// leandro_driguez: por qué el criterio es este?
	for len(indexList) < b {
		if j < b {
			indexList = append(indexList, j)
		}
		if i >= 0 {
			indexList = append(indexList, i)
		}
		i--
		j++
	}

	sl := &ShortList{}

	leftToAdd := num

	// Next we select alpha contacts and add them to the short list
	for leftToAdd > 0 && len(indexList) > 0 {
		index, indexList = indexList[0], indexList[1:]
		bucketContacts := len(rt.KBuckets[index])
		for i := 0; i < bucketContacts; i++ {
			ignored := false
			for j := 0; j < len(ignoredNodes); j++ {
				if bytes.Equal(rt.KBuckets[index][i].ID, ignoredNodes[j].ID) {
					ignored = true
				}
			}
			if !ignored {
				sl.AppendUnique([]*Node{&rt.KBuckets[index][i]})
				leftToAdd--
				if leftToAdd == 0 {
					break
				}
			}
		}
	}

	// TODO
	// leandro_driguez: no se está ordenando por el criterio de cercanía
	// respecto a un ID específico
	sort.Sort(sl)

	return sl
}
