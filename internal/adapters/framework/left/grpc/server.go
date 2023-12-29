package rpc

import (
	"log"
	"net"

	"github.com/Ayomide-Daniel/go-hex/internal/adapters/framework/left/grpc/pb"
	"github.com/Ayomide-Daniel/go-hex/internal/ports"

	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpcA Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to listen on port :9000, %v", err)
	}

	arithmeticServiceServer := grpcA
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to server gRPC server over port :9000: %v", err)
	}
}
