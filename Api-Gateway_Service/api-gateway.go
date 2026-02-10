package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	pb "github.com/M-Mongy/Distributed-System_Project/GRPC_Server"

	"google.golang.org/grpc"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *sql.DB

func main() {
	// اتصال بـ gRPC Service
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC: %v", err)
	}
	defer conn.Close()

	client := pb.UserServiceClient(conn)

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		CreateUserHandler(w, r, client)
	})

	log.Println("API Gateway running on :8080")
	http.ListenAndServe(":8080", nil)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request, client pb.UserServiceClient) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var u user
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 👈 هنا بنادي gRPC Service
	resp, err := client.CreateUser(context.Background(), &pb.UserRequest{
		Name:  u.Name,
		Email: u.Email,
	})
	if err != nil {
		http.Error(w, "gRPC call failed", http.StatusInternalServerError)
		return
	}

	// رجّع الـ response اللي جه من gRPC
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
