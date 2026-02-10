module github.com/M-Mongy/Distributed-System_Project/Api-Gateway_Service

go 1.25.6

require (
	github.com/M-Mongy/Distributed-System_Project/GRPC_Server v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.78.0
)

require (
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251029180050-ab9386a59fda // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace github.com/M-Mongy/Distributed-System_Project/GRPC_Server => ../GRPC_Server
