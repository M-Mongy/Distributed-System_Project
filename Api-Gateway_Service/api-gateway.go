package main

import (
	"context"
	//"database/sql"
	"encoding/json"
	"log"
	"net/http"

	inv "github.com/M-Mongy/Distributed-System_Project/GRPC_Server/Invoicer"

	"google.golang.org/grpc"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

//var db *sql.DB

func main() {
	conn, err := grpc.Dial("localhost:8089", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC: %v", err)
	}
	defer conn.Close()

	client := inv.NewInvoicerClient(conn)

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		CreateUserHandler(w, r, client)
	})

	log.Println("API Gateway running on :8080")
	http.ListenAndServe(":8080", nil)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request, client inv.InvoicerClient) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var u user
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := client.Create(context.Background(), &inv.CreateRequest{
		Amount:    &inv.Amount{Amount: 0, Currency: ""},
		From:      u.Name,
		To:        u.Email,
		VATnumber: "",
	})
	if err != nil {
		http.Error(w, "gRPC call failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
