package main

import (
	"log"
	"net"

	"github.com/dn-github/istio-k8s-mesh/ratings/pb"
	"github.com/dn-github/istio-k8s-mesh/ratings/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":3003")
	if err != nil {
		log.Fatalf("Failed to listen:  %v", err)
	}

	s := grpc.NewServer()
	service := server.NewRatingsImpl()
	pb.RegisterRatingServiceServer(s, service)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
