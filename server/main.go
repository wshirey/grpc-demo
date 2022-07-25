package main

import (
	"log"
	"net"

	"github.com/wshirey/grpc-demo/addresses"
	"github.com/wshirey/grpc-demo/health"
	"github.com/wshirey/grpc-demo/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	// open TCP port
	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("error listening on port: %v", err)
	}
	log.Println("listening on port 9090")

	// create new gRPC server
	s := grpc.NewServer()

	// instantiate services (with any dependencies)
	addressService := service.NewService()
	healthService := health.NewService()

	// register services with server
	addresses.RegisterAddressesServer(s, addressService)
	grpc_health_v1.RegisterHealthServer(s, healthService)

	// enables reflection endpoint
	reflection.Register(s)

	// start server
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
