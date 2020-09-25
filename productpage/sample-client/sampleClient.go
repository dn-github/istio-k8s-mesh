package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dn-github/istio-k8s-mesh/productpage/pb"
	"github.com/grab/async"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Enter QPS (0 for single run and exit): ")
	var qps int
	_, err := fmt.Scanf("%d", &qps)

	fmt.Println("Enter Concurrency:")
	var concurrency int
	_, err = fmt.Scanf("%d", &concurrency)

	conn, err := grpc.Dial("192.168.99.109:32053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer conn.Close()
	client := pb.NewProductPageServiceClient(conn)

	req := &pb.Book{
		Name: "The Book Thief",
	}

	sleepTime := time.Second / time.Duration(qps)

	tasks := make([]async.Task, concurrency)

	for i := 0; i < concurrency; i++ {
		tasks[i] = async.Repeat(context.Background(), sleepTime, func(ctx context.Context) (interface{}, error) {
			res, err := client.Product(ctx, req)
			if err != nil {
				log.Printf("error while calling gRPC: %v", err)
				return nil, err
			}
			log.Printf("Response from Service: %v", res)
			return nil, nil
		})
	}

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		async.CancelAll(tasks)
		os.Exit(0)
	}()

	async.WaitAll(tasks)
}
