package server

import (
	"context"
	"log"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc"

	pb "go-simple/api/proto/product/v1"
)

type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
	Product   pb.ProductServiceServer
}

func NewGRPCServer(p Params) *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterProductServiceServer(server, p.Product)

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			lis, err := net.Listen("tcp", ":9090")
			if err != nil {
				return err
			}
			go func() {
				log.Println("Starting gRPC server on :9090")
				if err := server.Serve(lis); err != nil {
					log.Fatalf("gRPC server failed: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping gRPC server")
			server.GracefulStop()
			return nil
		},
	})

	return server
}

func GRPCLifeCycle(server *grpc.Server) {
	// This function is intentionally empty.
	// The server starts via lifecycle hooks in server.NewGRPCServer.
}
