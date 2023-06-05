package core

import (
	"crypto/sha1"
	"encoding/base64"

	"github.com/science-engineering-art/gotify/peer/persistence"
	kademlia "github.com/science-engineering-art/kademlia-grpc/core"
)

type Peer struct {
	kademlia.FullNode
}

func NewPeer(mongoDbIP string, isBootstrapNode bool) *Peer {
	db := persistence.NewMongoDb("admin", "songs", mongoDbIP)
	newPeer := kademlia.NewFullNode("0.0.0.0", 8080, 32140, db, isBootstrapNode)

	return &Peer{*newPeer}
}

func (p *Peer) Store(data *[]byte) (string, error) {

	hash := sha1.Sum(*data)
	key := base64.RawStdEncoding.EncodeToString(hash[:])

	_, err := p.StoreValue(key, data)
	if err != nil {
		return "", nil
	}

	return key, nil
}
