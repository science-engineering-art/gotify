package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strconv"
	"strings"

	trackerCore "github.com/science-engineering-art/gotify/tracker/core"
	trackerUtils "github.com/science-engineering-art/gotify/tracker/utils"
	kademliaCore "github.com/science-engineering-art/kademlia-grpc/core"
	"github.com/science-engineering-art/kademlia-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/readline.v1"
)

var tracker *trackerCore.Tracker
var grpcServerAddress string

func main() {
	// Init CLI for using Full Node Methods
	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF, readline.ErrInterrupt
			break
		}
		input := strings.Split(line, " ")
		switch input[0] {
		case "node":
			if len(input) != 4 {
				displayHelp()
				continue
			}
			port, _ := strconv.Atoi(input[1])
			bPort, _ := strconv.Atoi(input[2])
			isB, _ := strconv.ParseBool(input[3])

			flag.Parse()

			ip := getIpFromHost()
			grpcServerAddress = ip + ":" + strconv.FormatInt(int64(port), 10)
			tracker, _ = trackerCore.NewTracker(ip, port, bPort, isB)
			go CreateGRPCServerFromFullNode(tracker.FullNode)

			fmt.Println("Node running at:", ip, ":", port)

		case "storeSong":
			if len(input) != 3 {
				displayHelp()
				continue
			}
			jsonMetadata := input[1]
			dataHash := input[2]
			ids := tracker.StoreSongMetadata(jsonMetadata, dataHash)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("Stored with IDs: ", ids)
		case "getSongs":
			if len(input) != 2 {
				displayHelp()
				continue
			}
			key := input[1]
			keyHash := trackerUtils.GetJsonMetadataKeyHash(key)
			songList := tracker.GetSongList(keyHash)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("The retrieved value is:", songList)
		case "dht":
			tracker.FullNode.PrintRoutingTable()
		}
	}
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

func displayHelp() {
	fmt.Println(`
help - This message
store <message> - Store a message on the network
get <key> - Get a message from the network
info - Display information about this node
	`)
}

func CreateGRPCServerFromFullNode(fullNode kademliaCore.FullNode) {
	grpcServer := grpc.NewServer()

	pb.RegisterFullNodeServer(grpcServer, &fullNode)
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
}
