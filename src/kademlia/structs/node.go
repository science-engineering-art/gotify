package structs

import (
	"bytes"
	"strings"
)

type Node struct {
	ID   []byte
	IP   string
	Port int
}

func (b Node) equal(other Node) bool {
	return bytes.Equal(b.ID, other.ID) &&
		strings.EqualFold(b.IP, other.IP) &&
		b.Port == other.Port
}
