package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	invoicer "github.com/M-Mongy/Distributed-System_Project/GRPC_Server/Invoicer"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

// myinvoicerServer implements the Invoicer gRPC service
type myinvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
	db *sql.DB
}

// Create handles CreateRequest and inserts a new user into the Postgres DB
func (s *myinvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(
		"INSERT INTO users(name, email) VALUES($1, $2)",
		req.From,
		req.To,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec(
		"INSERT INTO outbox_events(event_type, payload) VALUES($1, $2)",
		"user.created",
		fmt.Sprintf(`{"name":"%s","email":"%s"}`, req.From, req.To),
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &invoicer.CreateResponse{
		Pdf:  []byte("test"),
		Docx: []byte("test"),
	}, nil
}

func main() {
	// Connect to Postgres
	connStr := "postgres://admin:admin@localhost:5432/users_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Cannot open DB:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Cannot connect to Postgres:", err)
	}
	log.Println("Connected to Postgres")

	// Start listening on TCP port for gRPC
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()
	server := &myinvoicerServer{
		db: db, // pass DB to the server
	}
	invoicer.RegisterInvoicerServer(grpcServer, server)

	log.Println("gRPC Server running on :8089")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
