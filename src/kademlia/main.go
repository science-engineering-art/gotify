package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/science-engineering-art/spotify/src/kademlia/core"
	"github.com/science-engineering-art/spotify/src/kademlia/pb"
	"github.com/science-engineering-art/spotify/src/kademlia/structs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	fullNodeServer    core.FullNode
	grpcServerAddress string = "0.0.0.0:8080"
)

func main() {
	// Get params from command line
	ip := flag.String("ip", "0.0.0.0", "set ip address of the peer")
	port := flag.Int("port", 8080, "set port for TCP conection on GRPC server")
	flag.Parse()

	// Update params into global variables
	grpcServerAddress = *ip + ":" + strconv.FormatInt(int64(*port), 10)

	// Create a gRPC server full node
	go CreateFullNodeServer(ip, port)
}

func displayFlagHelp() {
	fmt.Println(`cli-example

Usage:
	cli-example --port [port]

Options:
	--help Show this screen.
	--ip=<ip> Local IP [default: 0.0.0.0]
	--port=[port] Local Port [default: 0]
	--bip=<ip> Bootstrap IP
	--bport=<port> Bootstrap Port
	--stun=<bool> Use STUN protocol for public addr discovery [default: true]`)
}

func displayHelp() {
	fmt.Println(`
help - This message
store <message> - Store a message on the network
get <key> - Get a message from the network
info - Display information about this node
	`)
}

func CreateFullNodeServer(ip *string, port *int) {
	fullNodeServer = *core.NewGrpcFullNodeServer(*ip, *port, &structs.Storage{})

	// Create gRPC Server
	grpcServer := grpc.NewServer()

	pb.RegisterFullNodeServer(grpcServer, &fullNodeServer)
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
