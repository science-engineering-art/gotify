package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
)

func GetHashesPowerSet(jsonData string) []string {
	var data map[string]interface{}
	hashesPowerSet := []string{}

	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
	}

	keys := getKeys(data)
	// fmt.Printf("\nKeys %v \n\n", keys)

	keysPowerSet := getKeysPowerSet(keys)

	for _, keysSet := range keysPowerSet {

		subJson := make(map[string]interface{})
		for _, key := range keysSet {
			subJson[key] = data[key]
		}
		jsonFromMap, _ := json.Marshal(subJson)
		jsonStrFromMap := string(jsonFromMap)
		if err != nil {
			fmt.Printf("could not marshal map: %s\n", err)
		}
		jsonHash := sha1.Sum([]byte(jsonStrFromMap))
		resultHash := jsonHash[:]
		hash := base64.RawStdEncoding.EncodeToString(resultHash)

		// fmt.Printf("\nKeySet: %v\nHash: %s\n\n", keysSet, hash)

		hashesPowerSet = append(hashesPowerSet, hash)
	}

	return hashesPowerSet
}

func getKeysPowerSet(keys []string) [][]string {
	powerSetSize := int(math.Pow(2, float64(len(keys))))
	result := make([][]string, 0, powerSetSize)

	var index int
	for index < powerSetSize {
		var subSet []string

		for j, elem := range keys {
			if index&(1<<uint(j)) > 0 {
				subSet = append(subSet, elem)
			}
		}
		result = append(result, subSet)
		index++
	}
	return result

}

func getKeys(data map[string]interface{}) []string {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	return keys
}
