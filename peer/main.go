package main

import (
	"fmt"
	"log"
	"net"

	"github.com/science-engineering-art/gotify/peer/persistence"
	kademlia "github.com/science-engineering-art/kademlia-grpc/core"
	"github.com/science-engineering-art/kademlia-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	db   = persistence.NewMongoDb("admin", "songs")
	ip   = "0.0.0.0"
	port = 8080
)

func main() {
	peer := kademlia.NewFullNode(ip, port, 32140, db, true)

	grpcServer := grpc.NewServer()

	pb.RegisterFullNodeServer(grpcServer, peer)
	reflection.Register(grpcServer)

	grpcAddr := fmt.Sprintf("%s:%d", ip, port)
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}
}
