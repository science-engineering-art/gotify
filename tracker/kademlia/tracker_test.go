package kademlia

import (
	"crypto/sha1"
	"fmt"
	"testing"
	"time"

	"github.com/jbenet/go-base58"
	"github.com/stretchr/testify/assert"
)

// Creates twenty DHTs and bootstraps each with the previous
// at the end all should know about each other
func TestBootstrapTwentyTrackers(t *testing.T) {
	done := make(chan bool)
	ip := "127.0.0.1"
	port := 3000
	bootIp := ip
	trackers := []*Tracker{}
	for i := 0; i < 20; i++ {
		tracker, _ := NewTracker(ip, port, bootIp, port-1)
		port++
		trackers = append(trackers, tracker)
		err := tracker.dht.CreateSocket()
		assert.NoError(t, err)
	}

	for _, tracker := range trackers {
		assert.Equal(t, 0, tracker.dht.NumNodes())
		go func(dht *DHT) {
			err := dht.Listen()
			assert.Equal(t, "closed", err.Error())
			done <- true
		}(tracker.dht)
		go func(dht *DHT) {
			err := dht.Bootstrap()
			assert.NoError(t, err)
		}(tracker.dht)
		time.Sleep(time.Millisecond * 200)
	}

	time.Sleep(time.Millisecond * 2000)

	for _, tracker := range trackers {
		assert.Equal(t, 19, tracker.dht.NumNodes())
		err := tracker.dht.Disconnect()
		assert.NoError(t, err)
		<-done
	}
}

func TestStoreAndGetJsonDataTwoNodes(t *testing.T) {
	done := make(chan bool)
	ip := "127.0.0.1"
	port := 3000

	tracker1, _ := NewTracker(ip, port, ip, port-1)
	tracker2, _ := NewTracker(ip, port+1, ip, port)

	err := tracker1.dht.CreateSocket()
	assert.NoError(t, err)

	err = tracker2.dht.CreateSocket()
	assert.NoError(t, err)

	go func() {
		err := tracker1.dht.Listen()
		assert.Equal(t, "closed", err.Error())
		done <- true
	}()

	go func() {
		err := tracker2.dht.Listen()
		assert.Equal(t, "closed", err.Error())
		done <- true
	}()

	time.Sleep(1 * time.Second)

	tracker2.dht.Bootstrap()

	jsonSongMetadata := `
	{
		"Artist": "Silvio Rodriguez",
		"Genre": "Trova",
		"Title": "Ojala"
	}
	`
	songData := []byte("adwdjwodijwidjdjqwdij")
	result := tracker1.StoreSongMetadata(jsonSongMetadata, songData)

	//Print Trackers Values
	fmt.Println("Tracker 2 Store:", tracker2.dht.store)

	time.Sleep(1 * time.Second)
	targetValue := []string{}
	songDataHash := sha1.Sum(songData)
	value := songDataHash[:]
	targetValue = append(targetValue, base58.Encode(value))
	for _, key := range result {
		value, found, _ := tracker1.dht.Get(key)
		fmt.Println(targetValue, value, found)
		//assert.Equal(t, true, stringSlicesEqual(targetValue, value))
	}

	err = tracker1.dht.Disconnect()
	assert.NoError(t, err)

	err = tracker2.dht.Disconnect()
	assert.NoError(t, err)

	<-done
	<-done
}

func stringSlicesEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, str1 := range slice1 {
		if str1 != slice2[i] {
			return false
		}
	}
	return true
}
