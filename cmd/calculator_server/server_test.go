package main

import (
	"context"
	"errors"
	"log"
	"net"
	"testing"

	pb "github.com/demoS11/calc-service/pkg/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context) (pb.CalculatorClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	pb.RegisterCalculatorServer(baseServer, NewServer())
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := pb.NewCalculatorClient(conn)

	return client, closer
}

func TestCalculatorService_Calculate(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out *pb.CalculateReply
		err error
	}

	tests := map[string]struct {
		in       *pb.CalculateRequest
		expected expectation
	}{
		"Add": {
			in: &pb.CalculateRequest{
				Number1:  1,
				Number2:  2,
				Operator: "add",
			},
			expected: expectation{
				out: &pb.CalculateReply{
					Result: 3,
				},
				err: nil,
			},
		},
		"Subtract": {
			in: &pb.CalculateRequest{
				Number1:  1,
				Number2:  2,
				Operator: "subtract",
			},
			expected: expectation{
				out: &pb.CalculateReply{
					Result: -1,
				},
				err: nil,
			},
		},
		"Multiply": {
			in: &pb.CalculateRequest{
				Number1:  1,
				Number2:  2,
				Operator: "multiply",
			},
			expected: expectation{
				out: &pb.CalculateReply{
					Result: 2,
				},
				err: nil,
			},
		},
		"Divide": {
			in: &pb.CalculateRequest{
				Number1:  1,
				Number2:  2,
				Operator: "divide",
			},
			expected: expectation{
				out: &pb.CalculateReply{
					Result: 0,
				},
				err: nil,
			},
		},
		"DivideByZero": {
			in: &pb.CalculateRequest{
				Number1:  1,
				Number2:  0,
				Operator: "divide",
			},
			expected: expectation{
				out: &pb.CalculateReply{},
				err: errors.New("rpc error: code = InvalidArgument desc = division by zero is not allowed"),
			},
		},
		"InvalidOperator": {
			in: &pb.CalculateRequest{
				Number1:  1,
				Number2:  0,
				Operator: "invalid",
			},
			expected: expectation{
				out: &pb.CalculateReply{},
				err: errors.New("rpc error: code = InvalidArgument desc = invalid operator: invalid"),
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			out, err := client.Calculate(ctx, tt.in)
			if err != nil {
				if tt.expected.err.Error() != err.Error() {
					t.Errorf("Err -> \nWant: %q\nGot: %q\n", tt.expected.err, err)
				}
			} else {
				if tt.expected.out.Result != out.Result {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out, out)
				}
			}

		})
	}
}
