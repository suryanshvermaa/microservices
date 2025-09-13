# Microservices using Go
This repository contains a collection of microservices built using Go. Each service is designed to be independent and can communicate with others over HTTP or gRPC.

# Protobuf generate command
```
protoc --go_out=. --go-grpc_out=. --proto_path=.  *.proto
```