package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/science-engineering-art/spotify/src/kademlia/core"
	"github.com/science-engineering-art/spotify/src/kademlia/structs"
	"gopkg.in/readline.v1"
)

var fullNode core.FullNode

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

			storage := structs.NewStorage()

			ip := getIpFromHost()

			fullNode = *core.NewFullNode(ip, port, bPort, storage, isB)

		case "store":
			if len(input) != 3 {
				displayHelp()
				continue
			}
			key := input[1]
			data := input[2]
			id, err := fullNode.StoreValue(key, data)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("Stored with ID: ", id)
		case "get":
			if len(input) != 2 {
				displayHelp()
				continue
			}
			key := input[1]
			value, err := fullNode.GetValue(key)
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("The retrived value is:", value)
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
