package rpc

import (
	"context"

	"github.com/Ayomide-Daniel/go-hex/internal/adapters/framework/left/grpc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpcA Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error

	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpcA.api.GetAddition(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpcA Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error

	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpcA.api.GetSubtraction(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpcA Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error

	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpcA.api.GetMultiplication(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpcA Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error

	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpcA.api.GetDivision(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}
