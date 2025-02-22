package main

import (
	"context"
	"log"
	"time"

	pb "github.com/durgadaspatil45/grpc-server-client-app/protobuf-spec"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Executing client grpc call...")
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to initialize grpc connection: %v", err)
	}
	defer conn.Close()
	client := pb.NewCustomerServiceClient(conn)

	req := &pb.CreateCustomerRequest{
		Customer: &pb.Customer{
			Name:    "ajaya",
			Email:   "ajaya@gmail.com",
			PhoneNo: "4334343434",
		},
	}

	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()
	res, err := client.CreateCustomer(ctx, req)
	if err != nil {
		log.Fatalf("error for calling the database :%v", err)
	}
	log.Printf("response from the server : %v", res.GetMessage())

}
