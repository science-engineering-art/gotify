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

	b58 "github.com/jbenet/go-base58"
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
			dataBytes := []byte(data)
			//fmt.Println("data bytes", dataBytes)
			err = sender.Send(&pb.Data{Init: 0, End: int32(len(dataBytes)), Buffer: dataBytes})
			if err != nil {
				fmt.Println(err.Error())
			}
			data_hash := sha1.Sum(dataBytes)
			id := data_hash[:]
			str := b58.Encode(id)
			fmt.Println("Stored ID: ", str, "Stored Data:", dataBytes)

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

		case "findnode":
			if len(input) != 4 {
				displayHelp()
				continue
			}
			ip := input[1]
			port, _ := strconv.Atoi(input[2])
			data := input[3]
			target := b58.Decode(data)

			client := GetFullNodeClient(&ip, &port)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			pbKBucket, err := client.FindNode(ctx, &pb.TargetID{ID: target})
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("The found nodes where: ", pbKBucket.Bucket)
			//fmt.Println("First node is: ", b58.Encode(pbKBucket.Bucket[0].ID))

		case "findvalue":
			if len(input) != 4 {
				displayHelp()
				continue
			}
			ip := input[1]
			port, _ := strconv.Atoi(input[2])
			data := input[3]
			target := b58.Decode(data)

			if len(target) == 0 {
				fmt.Println("Invalid target decoding.")
				continue
			}

			client := GetFullNodeClient(&ip, &port)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			receiver, err := client.FindValue(ctx, &pb.TargetID{ID: target})
			if err != nil {
				fmt.Println(err.Error())
			}

			buffer := []byte{}
			nearestNeighbors := []*pb.Node{}
			var init int32 = 0

			for {
				data, err := receiver.Recv()
				if data == nil {
					break
				}
				if len(data.KNeartestBuckets.Bucket) != 0 {
					nearestNeighbors = data.KNeartestBuckets.Bucket
					break
				}
				if init == data.Value.Init {
					buffer = append(buffer, data.Value.Buffer...)
					init = data.Value.End
				} else {
					fmt.Println(err.Error())
				}
			}
			foundValue := buffer
			if len(foundValue) == 0 {
				fmt.Println("Not found the requested value, this are alpha closest nodes:", nearestNeighbors)
			} else {
				fmt.Println("Found value:", foundValue)
			}
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

	id, _ := core.NewID(*ip, *port)
	log.Printf("start gRPC server on %s with id %s", listener.Addr().String(), b58.Encode(id))
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
