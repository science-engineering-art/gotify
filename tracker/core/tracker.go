package core

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/science-engineering-art/gotify/tracker/persistence"
	"github.com/science-engineering-art/gotify/tracker/utils"
	kademlia "github.com/science-engineering-art/kademlia-grpc/core"
)

type Tracker struct {
	kademlia.FullNode
}

func NewTracker(ip string, port int, bootPort int, isBoot bool) (*Tracker, error) {
	metadataStorage := persistence.NewMetadataStorage()
	fn := kademlia.NewFullNode(ip, port, bootPort, metadataStorage, isBoot)
	tracker := &Tracker{FullNode: *fn}
	return tracker, nil
}

func (t *Tracker) GetSongList(key string) []string {
	songList := []string{}

	flatArray, err := t.FullNode.GetValue(key, 0, 0)
	if err != nil {
		fmt.Println("Error when retrieving data:", err)
		return songList
	}

	formatedArray := getFormatedArray(flatArray)

	fmt.Printf("\nGetSongList(%s) => FormatedArray: %v\n\n", key, formatedArray)

	songList = getStringSliceFromByteArray(formatedArray)

	fmt.Printf("\nGetSongList(%s) => SongList: %v", key, songList)

	return songList
}

func (t *Tracker) StoreSongMetadata(jsonSongMetadata string, songDataHash string) []string {
	hashesPowerSet := utils.GetHashesPowerSet(jsonSongMetadata)
	fmt.Println("PowerSet", hashesPowerSet)
	valueFullJsonData := getValueFullJsonData(jsonSongMetadata, songDataHash)
	fmt.Println("ValueFullJsonData:", valueFullJsonData)
	for _, hash := range hashesPowerSet {
		// leandro_driguez: cambié el 2do parametro de StoreValue a []byte
		data := []byte(valueFullJsonData)
		fmt.Println("data", data)
		fmt.Println("&data", &data)
		fmt.Println("before crash", t.FullNode)
		_, err := t.FullNode.StoreValue(hash, &data)
		if err != nil {
			fmt.Println("Error when storing key:", hash, err)
		}
	}

	return hashesPowerSet
}

func getFormatedArray(flatArray []byte) [][]byte {
	result := [][]byte{}
	lenght := len(flatArray)

	for i := 0; i < lenght; {
		elemLen := int(binary.LittleEndian.Uint32(flatArray[i : i+4]))
		elem := flatArray[i+4 : i+4+elemLen]
		result = append(result, elem)
		i += 4 + elemLen
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
	var data map[string]string
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
