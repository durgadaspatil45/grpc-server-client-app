package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/durgadaspatil45/grpc-server-client-app/protobuf-spec"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedCustomerServiceServer
}

type Customer struct {
	Id      uint32
	Name    string
	Email   string
	PhoneNo string
}

func (s Server) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	cust := req.GetCustomer()
	dbReq := Customer{cust.Id, cust.Name, cust.Email, cust.PhoneNo}
	fmt.Printf("Customer: %v\n", dbReq)

	return &pb.CreateCustomerResponse{
		Message: "Added customer to the db",
	}, nil

}
func main() {
	log.Println("Starting the grpc server...")
	grpcServer := grpc.NewServer()
	pb.RegisterCustomerServiceServer(grpcServer, Server{})
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
