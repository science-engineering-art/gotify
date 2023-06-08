package main

import (
	"fmt"
	"log"
	"net"

	"github.com/science-engineering-art/gotify/peer/core"
	"github.com/science-engineering-art/gotify/peer/utils"
	"github.com/science-engineering-art/kademlia-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	ip   = utils.GetIpFromHost()
	port = 8080
)

func main() {
	peer := core.NewRedisPeer(ip, true)

	grpcServer := grpc.NewServer()

	pb.RegisterFullNodeServer(grpcServer, &peer.FullNode)
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
