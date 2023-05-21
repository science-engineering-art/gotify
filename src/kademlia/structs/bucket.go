package structs

import (
	"bytes"
	"strings"
)

type Bucket struct {
	ID   []byte
	IP   string
	Port int
}

func (b Bucket) equal(other Bucket) bool {
	return bytes.Equal(b.ID, other.ID) &&
		strings.EqualFold(b.IP, other.IP) &&
		b.Port == other.Port
}
