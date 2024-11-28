package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
	"time"
)

func sendMessages(rpcClient *rpc.Client, name string) {
	in := bufio.NewReader(os.Stdin)
	fmt.Printf("You are contnioue chat as %v:\n", name)
	for {

		line, _ := in.ReadString('\n')
		line = strings.TrimSpace(line)

		message := fmt.Sprintf("%s: %s", name, line)

		var messagesList []string
		err := rpcClient.Call("Listener.ChatRoom", message, &messagesList)
		if err != nil {
			log.Println("Error sending message:", err)
		}
	}
}

func receiveMessages(rpcClient *rpc.Client) {
	var lastMessageCount int

	for {

		var messagesList []string
		err := rpcClient.Call("Listener.GetMessages", "", &messagesList)
		if err != nil {
			log.Println("Error receiving messages:", err)
			time.Sleep(1 * time.Second)
			continue
		}

		if len(messagesList) > lastMessageCount {

			newMessages := messagesList[lastMessageCount:]
			for _, msg := range newMessages {
				fmt.Println(msg)
			}

			lastMessageCount = len(messagesList)
		}
	}
}

func main() {

	rpcClient, err := rpc.Dial("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	in := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name)

	go sendMessages(rpcClient, name)
	go receiveMessages(rpcClient)

	select {}
}
