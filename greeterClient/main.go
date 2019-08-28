// Package main implements a client for Greeter service.
package main

import (
	"bufio"
	"context"
	"fmt"
	pb "goGRPC/chat"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please type your name: ")
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1]
	for {
		fmt.Print("Please type your message: ")
		message, _ := reader.ReadString('\n')
		message = message[:len(message)-1]
		c := pb.NewGreeterClient(conn)
		// Contact the server and print out its response.
		r, err := c.SendMessage(context.TODO(), &pb.MessageRequest{Name: name, Message: message})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Println(r.Receive)
	}
}
