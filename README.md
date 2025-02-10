# gRPC with Go-Kit in Golang

## Overview
This project demonstrates how to build a gRPC service using Go-Kit in Golang. The example consists of a simple Greeter service that responds with a greeting message when a request is made.

## Features
- Implements gRPC with Go-Kit
- Follows Clean Architecture (Service, Endpoint, Transport layers)
- Supports gRPC client-server communication
- Easily extendable for additional functionalities

## Technologies Used
- **Golang** (Go 1.18+)
- **gRPC**
- **Go-Kit**
- **Protocol Buffers (proto3)**

## Project Structure
```
.
├── client
│   ├── main.go            # gRPC client implementation
├── server
│   ├── main.go            # gRPC server entry point
├── service
│   ├── service.go         # Business logic for Greeter service
├── endpoint
│   ├── endpoint.go        # Defines Go-Kit endpoints
├── transport
│   ├── transport.go       # gRPC transport layer for Go-Kit
├── proto
│   ├── helloworld.proto   # Protocol Buffer file
└── README.md
```

## Prerequisites
Ensure you have the following installed:
- [Go 1.18+](https://go.dev/dl/)
- [Protocol Buffers Compiler (protoc)](https://grpc.io/docs/protoc-installation/)
- gRPC and Go-Kit dependencies:
  ```sh
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  go get github.com/go-kit/kit
  ```

## Setup & Installation
### 1. Generate gRPC Code
Run the following command to generate Go code from the `.proto` file:
```sh
protoc --go_out=. --go-grpc_out=. proto/helloworld.proto
```

### 2. Run the gRPC Server
```sh
go run server/main.go
```

### 3. Run the gRPC Client
```sh
go run client/main.go
```

## Expected Output
- **Server Output:**
  ```
  gRPC server is listening on port 50051
  ```
- **Client Output:**
  ```
  Greeting: Hello, Go-Kit
  ```

## Code Flow Explanation
The code follows the Go-Kit architecture with three main layers:

1. **Service Layer (`service/service.go`)**
   - Implements the actual business logic.
   - Defines `GreeterService` interface and its implementation.
   
2. **Endpoint Layer (`endpoint/endpoint.go`)**
   - Converts incoming requests into structured endpoint requests.
   - Uses `Go-Kit` endpoints to abstract the business logic from the transport layer.
   
3. **Transport Layer (`transport/transport.go`)**
   - Implements the gRPC server.
   - Converts gRPC requests into endpoint requests and vice versa.
   - Uses Go-Kit transport handlers to decouple transport logic from business logic.

4. **Server Initialization (`server/main.go`)**
   - Creates an instance of the service.
   - Registers the gRPC transport layer.
   - Starts listening on port `50051`.

5. **Client (`client/main.go`)**
   - Connects to the gRPC server.
   - Sends a request to `SayHello` method.
   - Receives and prints the response from the server.

## Extending the Project
- Add middleware (logging, authentication, rate limiting)
- Integrate HTTP transport using Go-Kit
- Implement service discovery with Consul or etcd

## License
This project is open-source and available under the [MIT License](LICENSE).