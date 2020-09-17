package main

import (
	"log"
	"net"

	"github.com/dn-github/istio-k8s-mesh/productpage/pb"
	"github.com/dn-github/istio-k8s-mesh/productpage/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen:  %v", err)
	}

	s := grpc.NewServer()
	service := server.NewProductPageImpl()
	pb.RegisterProductPageServiceServer(s, service)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
