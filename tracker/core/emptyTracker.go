package core

import (
	"fmt"

	"github.com/science-engineering-art/gotify/tracker/persistence"
	"github.com/science-engineering-art/gotify/tracker/utils"
	kademlia "github.com/science-engineering-art/kademlia-grpc/core"
)

type EmptyTracker struct {
	kademlia.FullNode
}

func NewEmptyTracker(ip string, port int, bootPort int, isBoot bool) (*EmptyTracker, error) {
	storage := persistence.NewEmpty()
	fn := kademlia.NewFullNode(ip, port, bootPort, storage, isBoot)
	tracker := &EmptyTracker{FullNode: *fn}
	return tracker, nil
}

func (t *EmptyTracker) GetSongList(key string) []string {
	songList := []string{}

	flatArray, err := t.FullNode.GetValue(key, 0, 0)
	if err != nil {
		fmt.Println("Error when retrieving data:", err)
		return songList
	}

	formatedArray := getFormatedArray(flatArray)

	// fmt.Printf("\nGetSongList(%s) => FormatedArray: %v\n\n", key, formatedArray)

	songList = getStringSliceFromByteArray(formatedArray)

	// fmt.Printf("\nGetSongList(%s) => SongList: %v", key, songList)

	return songList
}

func (t *EmptyTracker) StoreSongMetadata(jsonSongMetadata string, songDataHash string) []string {
	hashesPowerSet := utils.GetHashesPowerSet(jsonSongMetadata)
	// fmt.Println("PowerSet", hashesPowerSet)
	valueFullJsonData := getValueFullJsonData(jsonSongMetadata, songDataHash)
	// fmt.Println("ValueFullJsonData:", valueFullJsonData)
	for _, hash := range hashesPowerSet {
		// leandro_driguez: cambié el 2do parametro de StoreValue a []byte
		data := []byte(valueFullJsonData)
		// fmt.Println("data", data)
		// fmt.Println("&data", &data)
		// fmt.Println("before crash", t.FullNode)
		_, err := t.FullNode.StoreValue(hash, &data)
		if err != nil {
			fmt.Println("Error when storing key:", hash, err)
		}
	}

	return hashesPowerSet
}
