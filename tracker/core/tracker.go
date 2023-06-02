package core

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/json"
	"fmt"

	"github.com/science-engineering-art/spotify/src/kademlia/core"
	"github.com/science-engineering-art/spotify/src/tracker/persistence"
	"github.com/science-engineering-art/spotify/src/tracker/utils"

	b58 "github.com/jbenet/go-base58"
)

type Tracker struct {
	fn core.FullNode
}

func NewTracker(ip string, port int, bootPort int, isBoot bool) (*Tracker, error) {
	metadataStorage := persistence.NewMetadataStorage()
	fn := core.NewFullNode(ip, port, bootPort, metadataStorage, isBoot)
	tracker := &Tracker{fn: *fn}
	return tracker, nil
}

func (t *Tracker) GetSongList(key string) []string {
	songList := []string{}

	flatArray, err := t.fn.GetValue(key)
	if err != nil {
		fmt.Println("Error when retrieving data:", err)
		return songList
	}

	formatedArray := getFormatedArray(flatArray)
	songList = getStringSliceFromByteArray(formatedArray)

	return songList
}

func (t *Tracker) StoreSongMetadata(jsonSongMetadata string, songDataHash string) []string {
	hashesPowerSet := utils.GetHashesPowerSet(jsonSongMetadata)

	for _, hash := range hashesPowerSet {
		id, err := t.fn.StoreValue(hash, songDataHash)
		if err != nil {
			fmt.Println("Error when storing key:", id, err)
		}
	}

	return hashesPowerSet
}

// newID generates a new random ID
func getNewID() ([]byte, error) {
	result := make([]byte, 20)
	_, err := rand.Read(result)
	return result, err
}

func getFormatedArray(flatArray []byte) [][20]byte {
	result := [][20]byte{}
	lenght := len(flatArray)
	currentElem := [20]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < lenght; i++ {
		pos := i % 10
		if pos == 0 {
			result = append(result, currentElem)
			currentElem := [20]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
			currentElem[pos] = flatArray[i]
		} else if pos < lenght-1 {
			currentElem[pos] = flatArray[i]
		} else {
			currentElem[pos] = flatArray[i]
			result = append(result, currentElem)
		}
	}

	return result
}

func getStringSliceFromByteArray(array [][20]byte) []string {
	result := []string{}
	for _, byteArray := range array {
		str := string(byteArray[:])
		result = append(result, str)
	}
	return result
}

func GetStringKeyFromRawJson(jsonSongMetadata string) string {
	subJson := make(map[string]interface{})
	jsonFromMap, err := json.Marshal(subJson)
	if err != nil {
		fmt.Printf("could not marshal map: %s\n", err)
	}
	jsonHash := sha1.Sum(jsonFromMap)
	resultHash := jsonHash[:]
	hash := b58.Encode(resultHash)
	return hash
}
