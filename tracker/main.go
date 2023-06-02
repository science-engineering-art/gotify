package main

import (
	"fmt"

	"github.com/science-engineering-art/spotify/src/tracker/utils"
)

func main() {
	jsonData := `
	{
		"intValue": 1234,
		"boolValue": true,
		"stringValue": "hello!"
	}
	`
	hashesPowerSet := utils.GetHashesPowerSet(jsonData)
	fmt.Println(hashesPowerSet)
}
