package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Listener int

func (l *Listener) GetLine(line []byte, ack *bool) error {
	fmt.Println(string(line))
	*ack = true
	return nil
}

var messagesList []string // Global shared list

func (l *Listener) ChatRoom(line string, messages *[]string) error {

	fmt.Println(line)

	messagesList = append(messagesList, line)

	*messages = messagesList

	fmt.Printf("FROM SERVER the NEW list is %v\n", messagesList)
	return nil
}

func main() {
	ddry, err := net.ResolveTCPAddr("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatal(err)
	}
	inbound, err := net.ListenTCP("tcp", ddry)
	if err != nil {
		log.Fatal(err)
	}
	listener := new(Listener)

	rpc.Register(listener)
	rpc.Accept(inbound)
}
