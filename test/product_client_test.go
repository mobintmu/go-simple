package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-simple/internal/config"
	"go-simple/internal/db/sqlc"
	"go-simple/internal/product/dto"
	"io"
	"net/http"
	"testing"
)

func TestProductsClient(t *testing.T) {
	// t.Parallel()
	WithTestServer(t, func() {
		cfg, err := config.NewConfig()
		if err != nil {
			t.Fatalf("Failed to load config: %v", err)
		}
		addr := fmt.Sprintf("http://%s:%d", cfg.HTTPAddress, cfg.HTTPPort)

		// First, create a product via admin API so it's available to the client
		var product dto.ProductResponse
		t.Run("Create Product (Admin)", func(t *testing.T) {
			productRequest := sqlc.CreateProductParams{
				Name:        "Client Visible Product",
				Description: "Visible to client",
				Price:       2000,
				IsActive:    true,
			}
			body, err := json.Marshal(productRequest)
			if err != nil {
				t.Fatalf("Failed to marshal product: %v", err)
			}
			resp, err := http.Post(addr+"/api/v1/admin/products", "application/json", bytes.NewBuffer(body))
			if err != nil {
				t.Fatalf("Failed to send request: %v", err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusCreated {
				t.Fatalf("Expected status 201 Created, got %d", resp.StatusCode)
			}
			if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
				t.Fatalf("Failed to decode response: %v", err)
			}
		})

		t.Run("List Products (Client)", func(t *testing.T) {
			resp, err := http.Get(addr + "/api/v1/products")
			if err != nil {
				t.Fatalf("Failed to send GET request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
				responseBody, _ := io.ReadAll(resp.Body)
				t.Logf("Response body: %s", string(responseBody))
				return
			}

			var products dto.ClientListProductsResponse
			if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
				t.Fatalf("Failed to decode response: %v", err)
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

		t.Run("Get Product By ID (Client)", func(t *testing.T) {
			resp, err := http.Get(fmt.Sprintf("%s/api/v1/products/%d", addr, product.ID))
			if err != nil {
				t.Fatalf("Failed to send GET request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
				responseBody, _ := io.ReadAll(resp.Body)
				t.Logf("Response body: %s", string(responseBody))
				return
			}

			var fetchedProduct dto.ProductResponse
			if err := json.NewDecoder(resp.Body).Decode(&fetchedProduct); err != nil {
				t.Fatalf("Failed to decode response: %v", err)
			}

			if fetchedProduct.ID != product.ID {
				t.Errorf("Expected product ID %d, got %d", product.ID, fetchedProduct.ID)
			}
			if fetchedProduct.Name != product.Name {
				t.Errorf("Expected product name %q, got %q", product.Name, fetchedProduct.Name)
			}
		})

		t.Run("Delete Product", func(t *testing.T) {
			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/v1/admin/products/%d", addr, product.ID), nil)
			if err != nil {
				t.Fatalf("Failed to create DELETE request: %v", err)
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Failed to send DELETE request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusNoContent {
				t.Errorf("Expected status 204 No Content, got %d", resp.StatusCode)
				responseBody, _ := io.ReadAll(resp.Body)
				t.Logf("Response body: %s", string(responseBody))
			}
		})
	})
}
