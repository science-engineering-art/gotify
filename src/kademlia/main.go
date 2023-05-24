package main

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/science-engineering-art/spotify/src/kademlia/core"
	"github.com/science-engineering-art/spotify/src/kademlia/pb"
	"github.com/science-engineering-art/spotify/src/kademlia/structs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"gopkg.in/readline.v1"
)

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
		case "help":
			displayHelp()
		case "peer":
			if len(input) != 3 {
				displayHelp()
				continue
			}
			ip := input[1]
			port, _ := strconv.Atoi(input[2])

			// Create a gRPC server full node
			go CreateFullNodeServer(&ip, &port)

		case "store":
			if len(input) != 4 {
				displayHelp()
				continue
			}
			ip := input[1]
			port, _ := strconv.Atoi(input[2])
			data := input[3]

			client := GetFullNodeClient(&ip, &port)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			sender, err := client.Store(ctx)
			if err != nil {
				fmt.Println(err.Error())
			}
			err = sender.Send(&pb.Data{Init: 0, End: int32(len(data)), Buffer: []byte(data)})
			if err != nil {
				fmt.Println(err.Error())
			}
			data_hash := sha1.Sum([]byte(data))
			id := string(data_hash[:])
			fmt.Printf("Stored ID: %s", id)

		case "ping":
			if len(input) != 5 {
				displayHelp()
				continue
			}
			ipSender := input[1]
			portSender, _ := strconv.Atoi(input[2])
			ipReceiver := input[3]
			portReceiver, _ := strconv.Atoi(input[4])
			client := GetFullNodeClient(&ipReceiver, &portReceiver)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			idSender, _ := core.NewID(ipSender, portSender)
			pbNode, err := client.Ping(ctx, &pb.Node{ID: idSender, IP: ipSender, Port: int32(portSender)})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("The requested node is alive at:", pbNode.IP, ":", pbNode.Port)
		}
	}
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
	grpcServerAddress := *ip + ":" + strconv.FormatInt(int64(*port), 10)
	fullNodeServer := *core.NewGrpcFullNodeServer(*ip, *port, structs.NewStorage())

	// Create gRPC Server
	grpcServer := grpc.NewServer()

	pb.RegisterFullNodeServer(grpcServer, &fullNodeServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", grpcServerAddress)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	log.Printf("start gRPC server %+v on %s", fullNodeServer.DHT.ID, listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}
}

func GetFullNodeClient(ip *string, port *int) pb.FullNodeClient {
	address := fmt.Sprintf("%s:%d", *ip, *port)
	conn, _ := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	client := pb.NewFullNodeClient(conn)
	return client
}
