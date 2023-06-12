package main

import (
	// "fmt"
	"log"
	"net"
	"strconv"

	trackerCore "github.com/science-engineering-art/gotify/tracker/core"
	"github.com/science-engineering-art/gotify/tracker/utils"
	"github.com/science-engineering-art/kademlia-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	ip   = utils.GetIpFromHost()
	port = 9090
)

func main() {
	// Init CLI for using Full Node Methods
	// rl, err := readline.New("> ")
	// if err != nil {
	// 	panic(err)
	// }
	// defer rl.Close()

	tracker, _ := trackerCore.NewTracker(ip, port, 42140, true)

	grpcServer := grpc.NewServer()

	pb.RegisterFullNodeServer(grpcServer, &tracker.FullNode)
	reflection.Register(grpcServer)

	grpcServerAddress := ip + ":" + strconv.FormatInt(int64(port), 10)
	listener, err := net.Listen("tcp", grpcServerAddress)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	//fmt.Println("Node running at:", ip, ":", port)

	// for {
	// 	line, err := rl.Readline()
	// 	if err != nil { // io.EOF, readline.ErrInterrupt
	// 		break
	// 	}
	// 	input := strings.Split(line, " ")
	// 	switch input[0] {
	// 	case "storeSong":
	// 		if len(input) != 3 {
	// 			displayHelp()
	// 			continue
	// 		}
	// 		jsonMetadata := input[1]
	// 		dataHash := input[2]
	// 		ids := tracker.StoreSongMetadata(jsonMetadata, dataHash)
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 		}
	// 		fmt.Println("Stored with IDs: ", ids)
	// 	case "getSongs":
	// 		if len(input) != 2 {
	// 			displayHelp()
	// 			continue
	// 		}
	// 		key := input[1]
	// 		keyHash := trackerUtils.GetJsonMetadataKeyHash(key)
	// 		songList := tracker.GetSongList(keyHash)
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 		}
	// 		fmt.Println("The retrieved value is:", songList)
	// 	case "dht":
	// 		tracker.FullNode.PrintRoutingTable()
	// 	}
	// }
}
