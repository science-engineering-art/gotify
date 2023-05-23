package structs

import (
	"fmt"
)

type Bucket struct {
	ID   string
	IP   string
	Port int
}

func (p *Bucket) init() {
	if p.ID == "" {
		p.ID = "9302923"
	}
}

func (p *Bucket) GetURL() string {
	return fmt.Sprintf("%s:%d", p.IP, p.Port)
}

func (p *Bucket) GetId() string {
	return p.ID
}
