package main

import (
	"context"

	pb "github.com/demoS11/calc-service/pkg/calculator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CalculatorService implements the calculator service.
type CalculatorService struct {
	pb.UnimplementedCalculatorServer
}

// NewServer creates a new instance of the CalculatorService to handle requests.
func NewServer() pb.CalculatorServer {
	t := &CalculatorService{}
	return t
}

// Calculate performs arithmetic calculations based on the provided operator and operands.
// Supported operators: "add", "subtract", "multiply", "divide".
func (s *CalculatorService) Calculate(ctx context.Context, req *pb.CalculateRequest) (*pb.CalculateReply, error) {
	var result int32

	switch req.GetOperator() {
	case "add":
		result = req.GetNumber1() + req.GetNumber2()
	case "subtract":
		result = req.GetNumber1() - req.GetNumber2()
	case "multiply":
		result = req.GetNumber1() * req.GetNumber2()
	case "divide":
		if req.GetNumber2() == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "division by zero is not allowed")
		}
		result = req.GetNumber1() / req.GetNumber2()
	default:
		return nil, status.Errorf(codes.InvalidArgument, "invalid operator: %s", req.GetOperator())
	}

	return &pb.CalculateReply{Result: result}, nil
}
