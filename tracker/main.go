package main

import "fmt"

func main() {
	jsonData := `
	{
		"intValue": 1234,
		"boolValue": true,
		"stringValue": "hello!"
	}
	`
	// hashesPowerSet := utils.GetHashesPowerSet(jsonData)
	// fmt.Println(hashesPowerSet)

	fmt.Println([]byte(jsonData), string([]byte(jsonData)))
}
