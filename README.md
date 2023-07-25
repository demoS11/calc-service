# calculator-service

This project implements a gRPC server and client with Go for performing simple math calculations. The server supports addition, subtraction, multiplication, and division operations for now.

This project utilizes Protocol Buffers (proto3) for defining the message formats and gRPC services.

## Table of Contents

- Introduction
- Installation
- Usage
- Unit Tests

## Requirements

- Go (1.14 or higher)
- gRPC (google.golang.org/grpc)
- Protobuf (github.com/golang/protobuf)

## Installation

1. **Clone the repository:**

   ```go
   git clone https://github.com/demoS11/calc-service.git
   cd calc-service
   ```

2. **Install Go**

   For installation see installation guide of go: [https://go.dev/doc/install](https://go.dev/doc/install)

3. **Install Protocol Buffer Compiler**

   For installation see installation guide of protoc3: [https://grpc.io/docs/protoc-installation/](https://grpc.io/docs/protoc-installation/)

4. **Install Go Plugins for the Protocol Compiler**

   For installation see installation guide of Go plugins for the protocol compiler: [https://grpc.io/docs/languages/go/quickstart/](https://grpc.io/docs/languages/go/quickstart/)

5. **Install the required dependencies**

   ```go
   go mod download
   ```

## Usage

1. **Build Server and Client**

   ```go
   make build
   ```

2. **Running the Server**

   To start the gRPC server, run the following command:

   ```go
   make run-server
   ```

   The server will start listening on port **`50051`** by default.

3. **Running the Client**

   To use the gRPC client for performing math calculations, use the **`client`** executable with the following options:

   Replace **`<operator>`** with the math operation to perform (**`add`**, **`subtract`**, **`multiply`**, or **`divide`**). Replace **`<number>`** with the operands for the respective operation.

   ### Examples

   - Addition:
     ```go
     ./client.out -method add -a 5 -b 3
     ```
   - Subtraction:
     ```go
     ./client.out -method subtract -a 10 -b 3
     ```
   - Multiplication:
     ```go
     ./client.out -method multiply -a 4 -b 7
     ```
   - Division:
     ```go
     ./client.out -method divide -a 15 -b 5
     ```

   The result of the calculation will be displayed on the console.

4. **Unit Tests**

   ```go
   make test
   ```
