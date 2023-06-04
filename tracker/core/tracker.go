package core

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/science-engineering-art/gotify/tracker/persistence"
	"github.com/science-engineering-art/gotify/tracker/utils"
	"github.com/science-engineering-art/kademlia-grpc/core"
)

type Tracker struct {
	FN core.FullNode
}

func NewTracker(ip string, port int, bootPort int, isBoot bool) (*Tracker, error) {
	metadataStorage := persistence.NewMetadataStorage()
	fn := core.NewFullNode(ip, port, bootPort, metadataStorage, isBoot)
	tracker := &Tracker{FN: *fn}
	return tracker, nil
}

func (t *Tracker) GetSongList(key string) []string {
	songList := []string{}

	flatArray, err := t.FN.GetValue(key)
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
	valueFullJsonData := getValueFullJsonData(jsonSongMetadata, songDataHash)

	for _, hash := range hashesPowerSet {
		id, err := t.FN.StoreValue(hash, valueFullJsonData)
		if err != nil {
			fmt.Println("Error when storing key:", id, err)
		}
	}

	return hashesPowerSet
}

func getFormatedArray(flatArray []byte) [][]byte {
	result := [][]byte{}
	lenght := len(flatArray)

	for i := 0; i < lenght; {
		elemLen := int32(binary.LittleEndian.Uint32(flatArray[i : i+4]))
		elem := flatArray[i+4 : i+4+int(elemLen)]
		result = append(result, elem)
		i += i + 4 + int(elemLen)
	}

	return result
}

func getStringSliceFromByteArray(array [][]byte) []string {
	result := []string{}
	for _, byteArray := range array {
		str := string(byteArray)
		result = append(result, str)
	}
	return result
}

func getValueFullJsonData(jsonSongMetadata string, songDataHash string) string {
	var data map[string]interface{}
	fullJsonData := ""

	err := json.Unmarshal([]byte(jsonSongMetadata), &data)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
	}
	data["datahash"] = songDataHash

	jsonData, _ := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal map: %s\n", err)
	}

	fullJsonData = string(jsonData)
	return fullJsonData
}
