package server

import (
	"context"
	"fmt"
	"log"

	"github.com/dn-github/istio-k8s-mesh/details/pb"
)

type detailsImpl struct{}

func NewDetailsImpl() *detailsImpl {
	return &detailsImpl{}
}

// Details ...
func (r *detailsImpl) Details(ctx context.Context, book *pb.Book) (*pb.Detail, error) {
	price := int64(100)
	genre := "suspense"
	logMessage := fmt.Sprintf("The price of %s is %d and the genre is %s", book.Name, price, genre)
	log.Println(logMessage)
	return &pb.Detail{
		Price: price,
		Genre: genre,
	}, nil
}
