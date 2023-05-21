package interfaces

import "github.com/science-engineering-art/spotify/src/kademlia/structs"

type KademliaProtocol interface {
	Ping(structs.Bucket) structs.Bucket
	Store(value *[]byte) error
	FindNode(target *[]byte) (kBucket *[]structs.Bucket)
	FindValue(infoHash *[]byte) (value *[]byte)
	// KNeartestBuckets(obj Identificable) (kBucket *[]s.Bucket)
}
