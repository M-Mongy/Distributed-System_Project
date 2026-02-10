package main

import (
	"context"
	"log"
	"net"

	invoicer "github.com/MohamedMongy917/Distributed-System_Project/github.com/MohamedMongy917/Distributed-System_Project/GRPC_Server/Invoicer"
	"google.golang.org/grpc"
)

type myinvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s *myinvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serverregister := grpc.NewServer()

	server := &myinvoicerServer{}

	invoicer.RegisterInvoicerServer(serverregister, server)

	err = serverregister.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
