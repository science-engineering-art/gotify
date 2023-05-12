package interfaces

import s "github.com/science-engineering-art/spotify/src/peer/core/structs"

type KademliaProtocol interface {
	Ping(peer Identificable)
	StoreValue(value *[]byte)
	FindNode(obj Identificable) (kBucket *[]s.Bucket)
	FindValue(infoHash Identificable) (value *[]byte)
	KNeartestBuckets(obj Identificable) (kBucket *[]s.Bucket)
}
