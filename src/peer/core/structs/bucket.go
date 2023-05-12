package structs

import (
	"fmt"
)

type Bucket struct {
	Id   string
	IP   string
	Port string
}

func (p *Bucket) GetURL() string {
	return fmt.Sprintf("%s:%s", p.IP, p.Port)
}

func (p *Bucket) GetId() string {
	return p.Id
}
