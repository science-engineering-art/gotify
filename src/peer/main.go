package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/science-engineering-art/spotify/src/peer/core/structs"
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

	go ListenEntryPoint()

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

func ListenEntryPoint() {
	listener, err := net.Listen("tcp", "127.0.0.1:7999")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			panic(err)
		}

		bucket := structs.Bucket{}
		err = json.Unmarshal(buffer[:n], &bucket)
		if err != nil {
			panic(err)
		}

		myBucket, _ := json.Marshal(structs.Bucket{
			ID:   "",
			IP:   "255.255.255.255",
			Port: 32110,
		})
		conn.Write(myBucket)
	}
}

func Broadcast() {
	conn, err := net.Dial("tcp", "127.0.255.255:7999")
	if err != nil {
		panic(err)
	}

	myBucket, _ := json.Marshal(structs.Bucket{
		ID:   "",
		IP:   "0.0.0.0",
		Port: 32110,
	})

	_, err = conn.Write(myBucket)
	if err != nil {
		panic(err)
	}

}
