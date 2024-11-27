package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	rpcClient, err := rpc.Dial("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	in := bufio.NewReader(os.Stdin)

	// Prompt user for their name
	fmt.Println("Enter your name:")
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name) // Clean the input

	for {
		fmt.Println("Enter your message:")
		line, _ := in.ReadString('\n')
		line = strings.TrimSpace(line) // Clean the input

		// Construct the message to send
		message := fmt.Sprintf("%s: %s", name, line)

		// Call the server and update the chat history
		var messagesList []string
		err := rpcClient.Call("Listener.ChatRoom", message, &messagesList)
		if err != nil {
			log.Fatal("RPC error:", err)
		}

		// Display the updated chat history
		fmt.Println("\nChat History:")
		for _, msg := range messagesList {
			fmt.Println(msg)
		}
	}
}
