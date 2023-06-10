package core

import (
	"crypto/sha1"
	"fmt"

	base58 "github.com/jbenet/go-base58"
	"github.com/science-engineering-art/gotify/peer/persistence"
	kademlia "github.com/science-engineering-art/kademlia-grpc/core"
)

type RedisPeer struct {
	kademlia.FullNode
}

func NewRedisPeer(ip string, isBootstrapNode bool) *RedisPeer {
	db := persistence.NewRedisDb(ip)
	newPeer := kademlia.NewFullNode(ip, 8080, 32140, db, isBootstrapNode)

	return &RedisPeer{*newPeer}
}

func (p *RedisPeer) Store(data *[]byte) (string, error) {

	hash := sha1.Sum(*data)
	key := base58.Encode(hash[:])

	fmt.Println("Before StoreValue()")
	_, err := p.StoreValue(key, data)
	if err != nil {
		return "", nil
	}
	fmt.Println("After StoreValue()")

	return key, nil
}
