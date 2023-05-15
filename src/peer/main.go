package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/science-engineering-art/spotify/src/peer/pb"
	"github.com/science-engineering-art/spotify/src/peer/rpc"
	"github.com/science-engineering-art/spotify/src/peer/services"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	ctx         context.Context
	mongoclient *mongo.Client

	songService    services.SongService
	songCollection *mongo.Collection

	grpcServerAddress string = "0.0.0.0:8080"
	mongoDbUri        string = "mongodb://user:password@db:27017/?maxPoolSize=20&w=majority"
)

func init() {
	// Get params from command line
	ip := flag.String("ip", "0.0.0.0", "set ip address of the peer")
	port := flag.Int64("port", 8080, "set port for TCP conection on GRPC server")
	mongoUri := flag.String("mongo", "mongodb://user:password@db:27017/?maxPoolSize=20&w=majority", "Mongo DB Uri instance")
	flag.Parse()

	// Update params into global variables
	grpcServerAddress = *ip + ":" + strconv.FormatInt(*port, 10)
	mongoDbUri = *mongoUri

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDbUri))
	if err != nil {
		log.Fatal(err)
	}

	ctx = context.TODO()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// Collections
	songCollection = client.Database("admin").Collection("songs")
	songService = services.NewSongService(songCollection, ctx)
}

func main() {
	defer mongoclient.Disconnect(ctx)

	songServer, err := rpc.NewGrpcSongServer(songCollection, songService)
	if err != nil {
		log.Fatal("cannot create grpc postServer: ", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterSongServiceServer(grpcServer, songServer)
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
