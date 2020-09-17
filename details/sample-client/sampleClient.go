package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/dn-github/istio-k8s-mesh/details/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:3002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}
	client := pb.NewDetailServiceClient(conn)

	req := &pb.Book{
		Name: "The Book Thief",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	res, err := client.Details(ctx, req)
	if err != nil {
		log.Fatalf("error while calling gRPC: %v", err)
	}
	log.Printf("Response from Service: %v", res)
}
