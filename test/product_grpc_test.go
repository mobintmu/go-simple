package test

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	productv1 "go-simple/api/proto/product/v1"
	"go-simple/internal/config"
	"go-simple/internal/product/dto"
)

func TestProductGRPCGetProductByID(t *testing.T) {
	WithHttpTestServer(t, func() {
		cfg, err := config.NewConfig()
		if err != nil {
			t.Fatalf("Failed to load config: %v", err)
		}
		addr := cfg.HTTPAddress + ":" + strconv.Itoa(cfg.GRPCPort)
		conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		assert.NoError(t, err)
		defer conn.Close()
		client := productv1.NewProductServiceClient(conn)
		var product dto.ProductResponse
		grpcCreateProduct(t, &product, client)
		grpcGetProductByID(t, product, client)
		grpcListProducts(t, product, client)
		grpcUpdateProduct(t, product, client)
		grpcDeleteProduct(t, product, client)
	})
}

func grpcCreateProduct(t *testing.T, product *dto.ProductResponse, client productv1.ProductServiceClient) {
	t.Run("create Product (gRPC)", func(t *testing.T) {
		req := &productv1.CreateProductRequest{
			Name:        "Test Product gRPC",
			Description: "A product created during gRPC testing",
			Price:       30,
		}
		resp, err := client.CreateProduct(context.Background(), req)
		if err != nil {
			t.Fatalf("gRPC CreateProduct failed: %v", err)
		}
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Id)
		product.ID = resp.Id
		product.Name = req.Name
		product.Description = req.Description
		product.Price = req.Price
		t.Logf("gRPC Created product: %+v", product)
	})
}

func grpcGetProductByID(t *testing.T, product dto.ProductResponse, client productv1.ProductServiceClient) {
	t.Run("Get Product by ID (gRPC)", func(t *testing.T) {
		req := &productv1.ProductRequest{
			Id: product.ID,
		}
		resp, err := client.GetProductByID(context.Background(), req)
		if err != nil {
			t.Fatalf("gRPC GetProductByID failed: %v", err)
		}
		assert.NotNil(t, resp)
		assert.Equal(t, product.ID, resp.Id)
		assert.Equal(t, product.Name, resp.Name)
		assert.Equal(t, product.Description, resp.Description)
		assert.Equal(t, product.Price, resp.Price)
		t.Logf("gRPC Retrieved product by ID: %+v", resp)
	})
}

func grpcListProducts(t *testing.T, product dto.ProductResponse, client productv1.ProductServiceClient) {
	t.Run("List Products (gRPC)", func(t *testing.T) {
		resp, err := client.ListProducts(context.Background(), nil)
		if err != nil {
			t.Fatalf("gRPC ListProducts failed: %v", err)
		}
		assert.NotNil(t, resp)
		found := false
		for _, p := range resp.Products {
			if p.Id == product.ID {
				found = true
				t.Logf("gRPC Found product in list: %+v", p)
				break
			}
		}
		if !found {
			t.Errorf("gRPC Product not found in list")
		}
	})
}

func grpcUpdateProduct(t *testing.T, product dto.ProductResponse, client productv1.ProductServiceClient) {
	t.Run("Update Product (gRPC)", func(t *testing.T) {
		req := &productv1.UpdateProductRequest{
			Id:          product.ID,
			Name:        "Updated gRPC Product",
			Description: "An updated product during gRPC testing",
			Price:       50,
			IsActive:    true,
		}
		resp, err := client.UpdateProduct(context.Background(), req)
		if err != nil {
			t.Fatalf("gRPC UpdateProduct failed: %v", err)
		}
		assert.NotNil(t, resp)
		assert.Equal(t, req.Id, resp.Id)
		assert.Equal(t, req.Name, resp.Name)
		assert.Equal(t, req.Description, resp.Description)
		assert.Equal(t, req.Price, resp.Price)
		t.Logf("gRPC Updated product: %+v", resp)
	})
}

func grpcDeleteProduct(t *testing.T, product dto.ProductResponse, client productv1.ProductServiceClient) {
	t.Run("Delete Product (gRPC)", func(t *testing.T) {
		req := &productv1.DeleteProductRequest{
			Id: product.ID,
		}
		resp, err := client.DeleteProduct(context.Background(), req)
		if err != nil {
			t.Fatalf("gRPC DeleteProduct failed: %v", err)
		}
		assert.NotNil(t, resp)
		assert.Equal(t, product.ID, resp.Id)
		t.Logf("gRPC Deleted product ID: %d", resp.Id)
	})
}
