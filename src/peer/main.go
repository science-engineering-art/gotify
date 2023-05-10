package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/science-engineering-art/spotify/src/peer/api"
	"github.com/science-engineering-art/spotify/src/peer/services"
	pb "github.com/science-engineering-art/spotify/src/rpc/songs"

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
	DBUri             string = "mongodb://user:password@db:27017/?maxPoolSize=20&w=majority"
)

func init() {
	ctx = context.TODO()

	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI(DBUri)
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// Collections
	songCollection = mongoclient.Database("gotify").Collection("songs")
	songService = services.NewSongService(songCollection, ctx)
}

func main() {
	defer mongoclient.Disconnect(ctx)

	songServer, err := api.NewGrpcSongServer(songCollection, songService)
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
