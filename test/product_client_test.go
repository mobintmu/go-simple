package test

import (
	"encoding/json"
	"fmt"
	"go-simple/internal/auth"
	"go-simple/internal/config"
	"go-simple/internal/product/dto"
	"io"
	"net/http"
	"testing"
)

func TestProductsClient(t *testing.T) {
	// t.Parallel()
	WithHttpTestServer(t, func() {
		cfg, err := config.NewConfig()
		if err != nil {
			t.Fatalf("Failed to load config: %v", err)
		}
		addr := fmt.Sprintf("http://%s:%d", cfg.HTTPAddress, cfg.HTTPPort)

		// First, create a product via admin API so it's available to the client
		var product dto.ProductResponse
		token, err := auth.GenerateToken(cfg, "admin-123")
		if err != nil {
			t.Fatalf("Failed to generate token: %v", err)
		}
		adminCreateProduct(t, &product, addr+"/api/v1/admin/products", token)
		clientListProducts(t, product, addr)
		clientGetProductByID(t, product, addr)
		adminDeleteProduct(t, product, addr+"/api/v1/admin/products", token)
	})
}

func clientListProducts(t *testing.T, product dto.ProductResponse, addr string) {
	t.Run("List Products (Client)", func(t *testing.T) {
		resp, err := http.Get(addr + "/api/v1/products")
		if err != nil {
			t.Fatalf("Failed to send GET request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
			responseBody, _ := io.ReadAll(resp.Body)
			t.Logf(ResponseBodyMessage, string(responseBody))
			return
		}

		var products dto.ClientListProductsResponse
		if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
			t.Fatalf(FailedToDecodeMessage, err)
		}

		found := false
		for _, p := range products {
			if p.ID == product.ID {
				found = true
				t.Logf("Found product in client list: %+v", p)
				break
			}
		}
		if !found {
			t.Errorf("Product not found in client list")
		}
	})
}

func clientGetProductByID(t *testing.T, product dto.ProductResponse, addr string) {
	t.Run("Get Product By ID (Client)", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/api/v1/products/%d", addr, product.ID))
		if err != nil {
			t.Fatalf("Failed to send GET request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
			responseBody, _ := io.ReadAll(resp.Body)
			t.Logf(ResponseBodyMessage, string(responseBody))
			return
		}

		var fetchedProduct dto.ProductResponse
		if err := json.NewDecoder(resp.Body).Decode(&fetchedProduct); err != nil {
			t.Fatalf(FailedToDecodeMessage, err)
		}

		if fetchedProduct.ID != product.ID {
			t.Errorf("Expected product ID %d, got %d", product.ID, fetchedProduct.ID)
		}
		if fetchedProduct.Name != product.Name {
			t.Errorf("Expected product name %q, got %q", product.Name, fetchedProduct.Name)
		}
	})
}
