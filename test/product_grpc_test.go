package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	productv1 "go-simple/api/proto/product/v1"
	"go-simple/internal/config"
)

func TestProductGRPCGetProductByID(t *testing.T) {
	WithHttpTestServer(t, func() {
		cfg, err := config.NewConfig()
		if err != nil {
			t.Fatalf("Failed to load config: %v", err)
		}
		_ = cfg
		addr := fmt.Sprintf("localhost:9090")
		conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		assert.NoError(t, err)
		defer conn.Close()

		client := productv1.NewProductServiceClient(conn)

		// Replace with a valid seeded product ID
		req := &productv1.ProductRequest{Id: 4}
		resp, err := client.GetProductByID(context.Background(), req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, int32(4), resp.Id)
		assert.NotEmpty(t, resp.Name)
		assert.NotEmpty(t, resp.Description)
		assert.True(t, resp.Price > 0)
	})
}
