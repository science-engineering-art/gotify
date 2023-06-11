package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strconv"
	"strings"

	trackerCore "github.com/science-engineering-art/gotify/tracker/core"
	kademliaCore "github.com/science-engineering-art/kademlia-grpc/core"
	"github.com/science-engineering-art/kademlia-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var tracker *trackerCore.Tracker
var grpcServerAddress string

func main() {
	// Init CLI for using Full Node Methods
	// rl, err := readline.New("> ")
	// if err != nil {
	// 	panic(err)
	// }
	// defer rl.Close()

	port := 8081
	bPort := 5555
	isB := true

	ip := getIpFromHost()
	grpcServerAddress = ip + ":" + strconv.FormatInt(int64(port), 10)
	tracker, _ = trackerCore.NewTracker(ip, port, bPort, isB)
	grpcServer := grpc.NewServer()

	pb.RegisterFullNodeServer(grpcServer, &tracker.FullNode)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", grpcServerAddress)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	fmt.Println("Node running at:", ip, ":", port)

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

func getIpFromHost() string {
	cmd := exec.Command("hostname", "-i")
	var out strings.Builder
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running docker inspect:", err)
		return ""
	}
	ip := strings.TrimSpace(out.String())
	return ip
}

func CreateGRPCServerFromFullNode(fullNode kademliaCore.FullNode) {

}
