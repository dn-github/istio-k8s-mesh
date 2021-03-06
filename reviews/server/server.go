package server

import (
	"context"
	"fmt"
	"log"

	ratingPb "github.com/dn-github/istio-k8s-mesh/ratings/pb"
	"github.com/dn-github/istio-k8s-mesh/reviews/client"
	reviewPb "github.com/dn-github/istio-k8s-mesh/reviews/pb"
	"google.golang.org/grpc"
)

type reviewsImpl struct {
	ratingClient ratingPb.RatingServiceClient
}

func NewReviewsImpl() *reviewsImpl {
	conn, err := grpc.Dial("ratings:3003", grpc.WithInsecure())
	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}
	return &reviewsImpl{
		ratingClient: ratingPb.NewRatingServiceClient(conn),
	}
}

// Reviews v1 does nto call rating service
//func (r *reviewsImpl) Reviews(ctx context.Context, book *reviewPb.Book) (*reviewPb.Review, error) {
//	review := "I am version 1 and " + book.Name + " is a wonderful book"
//	log.Println(review)
//	return &reviewPb.Review{
//		Review: review,
//		Rating: 0,
//	}, nil
//}

// Reviews v2 calls rating service and shows ratings from 1 to 5
//func (r *reviewsImpl) Reviews(ctx context.Context, book *reviewPb.Book) (*reviewPb.Review, error) {
//	rating, err := client.GetRating(r.ratingClient, book.Name)
//	if err != nil {
//		log.Fatalf(err.Error())
//		rating = 0
//	}
//	review := fmt.Sprintf("I am version 2 and %s is a nice and wonderful book with rating %d", book.Name, rating)
//	log.Println(review)
//	return &reviewPb.Review{
//		Review: "I am version 2 and this is a nice and wonderful book",
//		Rating: rating,
//	}, nil
//}

// Reviews v3 calls rating service and shows ratings from 6 to 10
func (r *reviewsImpl) Reviews(ctx context.Context, book *reviewPb.Book) (*reviewPb.Review, error) {
	rating, err := client.GetRating(r.ratingClient, book.Name)
	if err != nil {
		log.Fatalf(err.Error())
		rating = 0
	}
	review := fmt.Sprintf("I am version 3 and %s is an extremely wonderful book with rating %d", book.Name, rating)
	log.Println(review)
	return &reviewPb.Review{
		Review: "I am version 3 and this is an extremely wonderful book",
		Rating: rating + 5,
	}, nil
}
