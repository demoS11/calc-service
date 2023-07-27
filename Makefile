# Go binary names for the server and client
SERVER_BINARY := server.out
CLIENT_BINARY := client.out

.PHONY: build test run-server run-client clean

# Default target when running `make` without any specific target
all: build

# Build the server and client binaries
build:
	@echo "Building server and client binaries..."
	@go build -o $(SERVER_BINARY) ./cmd/calculator_server
	@go build -o $(CLIENT_BINARY) ./cmd/calculator_client

# Run the gRPC server
run-server:
	@echo "Starting gRPC server..."
	@./$(SERVER_BINARY)

# Run the client application (modify arguments as needed)
run-client:
	@echo "Running the client application..."
	@./$(CLIENT_BINARY) -method add -a 1 -b 2

# Display help for client system
client-help:
	@echo "Client Usage:"
	@echo "  ./client.out -method <operator> -a <number> -b <number>"
	@echo ""
	@echo "  Options:"
	@echo "    -method string   Method to execute (add, subtract, multiply, divide)"
	@echo "    -a int           First operand"
	@echo "    -b int           Second operand"

# Run unit tests
test:
	@echo "Running unit tests..."
	@go test -v ./...

# Clean build artifacts and temporary files
clean:
	@echo "Cleaning up..."
	@rm -f $(SERVER_BINARY) $(CLIENT_BINARY)

# Help target to display available targets and their descriptions
help:
	@echo "Available targets:"
	@echo "  build       - Build the server and client binaries"
	@echo "  run-server  - Start the gRPC server"
	@echo "  run-client  - Run the client application"
	@echo "  test        - Run unit tests"
	@echo "  clean       - Clean up build artifacts and temporary files"

# Run linting using golangci-lint
lint:
	golangci-lint run ./...

# By default, show the help message
default: help
