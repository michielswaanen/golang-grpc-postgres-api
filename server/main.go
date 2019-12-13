package main

import (
	"../services"
	"./database"
	"./service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	// Create a new server listener
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err.Error())
	}

	database.NewConnection()

	// Create a new services server
	gRPCServer := grpc.NewServer()

	// Register a new Service Server
	services.RegisterAccountServiceServer(gRPCServer, service.InitializeAccountServer())

	// Serialize & Deserialize data
	reflection.Register(gRPCServer)

	// Start server listening
	if e := gRPCServer.Serve(listener); e != nil {
		panic(e.Error())
		return
	}
}
