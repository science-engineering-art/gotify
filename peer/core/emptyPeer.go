package core

import (
	"crypto/sha1"

	base58 "github.com/jbenet/go-base58"
	"github.com/science-engineering-art/gotify/peer/persistence"
	kademlia "github.com/science-engineering-art/kademlia-grpc/core"
)

type EmptyPeer struct {
	kademlia.FullNode
}

func NewEmptyPeer(ip string, isBootstrapNode bool) *EmptyPeer {
	db := persistence.NewEmpty()
	newPeer := kademlia.NewFullNode(ip, 8080, 32140, db, isBootstrapNode)

	return &EmptyPeer{*newPeer}
}

func (p *EmptyPeer) Store(data *[]byte) (string, error) {

	hash := sha1.Sum(*data)
	key := base58.Encode(hash[:])

	//fmt.Println("Before StoreValue()")
	_, err := p.StoreValue(key, data)
	if err != nil {
		return "", nil
	}
	//fmt.Println("After StoreValue()")

	return key, nil
}
