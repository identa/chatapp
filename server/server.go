package main

import (
	"grpcChatServer/chatserver"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "5000"
	}

	listen, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		log.Fatalf("Could not listen @ %v :: %v", Port, err)
	}
	log.Println("Listening @ : " + Port)

	grpcServer := grpc.NewServer()

	//register ChatService
	cs := chatserver.ChatServer{}
	chatserver.RegisterServicesServer(grpcServer, &cs)

	err = grpcServer.Serve(listen)

	if err != nil {
		log.Fatalf("Failed to start")
	}
}
