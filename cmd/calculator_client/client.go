package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	pb "github.com/demoS11/calc-service/pkg/calculator"
	"google.golang.org/grpc"
)

// Operation represents an arithmetic operation to be performed.
type Operation struct {
	Number1  int32
	Number2  int32
	Operator string
}

// validOperators stores the valid operator names.
var validOperators = map[string]bool{
	"add":      true,
	"subtract": true,
	"multiply": true,
	"divide":   true,
}

// isValidOperator checks if the given operator is valid.
func isValidOperator(operator string) bool {
	_, ok := validOperators[operator]
	return ok
}

// isFlagPassed checks if the flag is given or not given.
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

// parseArgs parses the command-line arguments and returns the operation to be performed.
func parseArgs() (Operation, error) {
	method := flag.String("method", "", "Method to execute (add, subtract, multiply, divide)")
	a := flag.Int("a", 0, "First operand")
	b := flag.Int("b", 0, "Second operand")
	flag.Parse()

	operator := strings.ToLower(*method)
	if !isFlagPassed("method") || !isValidOperator(operator) {
		return Operation{}, fmt.Errorf("Invalid or missing method. Use one of: add, subtract, multiply, divide")
	}

	if !isFlagPassed("a") || !isFlagPassed("b") {
		return Operation{}, fmt.Errorf("Missing operand. Both operands (a and b) are required")
	}

	return Operation{
		Number1:  int32(*a),
		Number2:  int32(*b),
		Operator: operator,
	}, nil
}

// calculate performs the arithmetic calculation using the gRPC client.
func calculate(conn *grpc.ClientConn, operation Operation) (int32, error) {
	req := &pb.CalculateRequest{
		Number1:  operation.Number1,
		Number2:  operation.Number2,
		Operator: operation.Operator,
	}

	client := pb.NewCalculatorClient(conn)

	result, err := client.Calculate(context.Background(), req)
	if err != nil {
		return 0, fmt.Errorf("Error calculating: %v", err)
	}

	return result.GetResult(), nil
}

// Parse command-line arguments and obtain the arithmetic operation to be performed
// Create a gRPC connection to the server running on localhost:50051
// Perform the arithmetic calculation using the gRPC client and the provided operation
// Print the result of the calculation to the console
func main() {
	operation, err := parseArgs()
	if err != nil {
		log.Fatalf("Error parsing arguments: %v", err)
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	result, err := calculate(conn, operation)
	if err != nil {
		log.Fatalf("Error calculating: %v", err)
	}

	fmt.Println(result)
}
