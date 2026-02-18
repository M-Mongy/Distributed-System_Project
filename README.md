## 📦 Distributed-System_Project

**Distributed-System_Project** is a Go-based distributed system example that demonstrates a simple microservice architecture built with an API Gateway, a gRPC server, and a consumer service, all orchestrated using Docker Compose.

This repository provides a solid foundation for learning and building distributed services in Go, with clear service separation, inter-service communication via gRPC, and database integration.

### 🧩 Project Structure

- **Api-Gateway_Service** – Serves as the entry point for client requests, handling HTTP/REST traffic and routing to backend services.
- **GRPC_Server** – Implements core business logic exposed over gRPC for efficient inter-service and client communication.
- **consumer** – Represents a background worker or service that consumes data or messages from other components.
- **docker-compose.yml** – Defines the development environment including PostgreSQL as the primary database for persistence.
- **go.mod / go.sum** – Dependency management for the Go modules used across the services.

### 🚀 Features

- Implemented entirely in **Go**, showcasing idiomatic patterns for microservices and distributed systems.
- Clear separation of services to support independent development, testing, and deployment.
- **Docker Compose** support for easy local environment setup with the database.
- Integrated **PostgreSQL** for persistent storage.

### 🛠️ Use Cases

This project can be used as:

- A learning resource for building distributed systems in Go.
- A template for microservice architecture with API gateway and gRPC communication.
- A starting point for extending with message queues, authentication, or advanced orchestration.

