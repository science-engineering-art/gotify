package main

import (
	"fmt"

	"github.com/science-engineering-art/spotify/src/tracker/kademlia"
)

func main() {
	jsonData := `
	{
		"intValue": 1234,
		"boolValue": true,
		"stringValue": "hello!"
	}
	`
	hashesPowerSet := kademlia.GetHashesPowerSet(jsonData)
	fmt.Println(hashesPowerSet)
}
