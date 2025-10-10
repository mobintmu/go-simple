package server

import (
	"context"
	"log"
	"net"
	"strconv"

	"go.uber.org/fx"
	"google.golang.org/grpc"

	pb "go-simple/api/proto/product/v1"
	"go-simple/internal/config"
)

type Params struct {
	fx.In
	Lifecycle fx.Lifecycle
	Product   pb.ProductServiceServer
	Config    *config.Config
}

func CreateGRPCServer(p Params) *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterProductServiceServer(server, p.Product)
	return server
}

func StartGRPCServer(p Params, server *grpc.Server) {
	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			lis, err := net.Listen("tcp", ":"+strconv.Itoa(p.Config.GRPCPort))
			if err != nil {
				return err
			}
			go func() {
				log.Println("🚀 Starting gRPC server on :" + strconv.Itoa(p.Config.GRPCPort))
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
}

func GRPCLifeCycle(server *grpc.Server) {
	// This function is intentionally empty.
	// The server starts via lifecycle hooks in server.NewGRPCServer.
}
