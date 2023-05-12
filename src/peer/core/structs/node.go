package structs

import (
	"fmt"
)

type Node struct {
	Id   string
	IP   string
	Port string
}

func (p *Node) GetURL() string {
	return fmt.Sprintf("%s:%s", p.IP, p.Port)
}

func (p *Node) GetId() string {
	return p.Id
}
