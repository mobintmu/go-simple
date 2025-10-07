package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-simple/internal/config"
	"go-simple/internal/product/dto"
	"io"
	"net/http"
	"testing"
)

func TestProductsAdmin(t *testing.T) {
	// t.Parallel()
	WithTestServer(t, func() {
		cfg, err := config.NewConfig()
		if err != nil {
			t.Fatalf("Failed to load config: %v", err)
		}
		addr := fmt.Sprintf("http://%s:%d", cfg.HTTPAddress, cfg.HTTPPort)
		addr += "/api/v1/admin/products"
		product := dto.ProductResponse{}
		adminCreateProduct(t, &product, addr)
		adminListProduct(t, product, addr)
		adminGetProductByID(t, product, addr)
		adminUpdateProduct(t, product, addr)
		adminDeleteProduct(t, product, addr)
		adminVerifyProductCreated(t, product, addr)
	})
}

func adminCreateProduct(t *testing.T, product *dto.ProductResponse, addr string) {
	t.Run("Create Product", func(t *testing.T) {
		productRequest := dto.AdminCreateProductRequest{
			Name:        "Test Product",
			Description: "This is a test product",
			Price:       1000,
			IsActive:    true,
		}
		body, err := json.Marshal(productRequest)
		if err != nil {
			t.Fatalf("Failed to marshal product: %v", err)
		}
		resp, err := http.Post(addr, "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Expected status 201 Created, got %d", resp.StatusCode)
			// Optional: Print response body for debugging
			responseBody, _ := io.ReadAll(resp.Body)
			t.Logf(ResponseBodyMessage, string(responseBody))
		}
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&product); err != nil {
			t.Fatalf(FailedToDecodeMessage, err)
		}
	})
}

func adminListProduct(t *testing.T, product dto.ProductResponse, addr string) {
	t.Run("List Products", func(t *testing.T) {
		resp, err := http.Get(addr)
		if err != nil {
			t.Fatalf(FailedToSendGetMessage, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf(ExpectedStatus200OKGotMessage, resp.StatusCode)
			responseBody, _ := io.ReadAll(resp.Body)
			t.Logf(ResponseBodyMessage, string(responseBody))
			return
		}

		var products dto.ClientListProductsResponse
		if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
			t.Fatalf(FailedToDecodeMessage, err)
		}

		if len(products) == 0 {
			t.Errorf("Expected at least one product, got 0")
		}
		findCreatedProduct := false
		for index, p := range products {
			if p.ID == product.ID {
				findCreatedProduct = true
				t.Logf("Found created product at index %d: %+v", index, p)
				break
			}
		}
		if !findCreatedProduct {
			t.Errorf("Created product not found in the list")
		}
	})
}

func adminGetProductByID(t *testing.T, product dto.ProductResponse, addr string) {
	t.Run("Get Product By ID", func(t *testing.T) {
		url := fmt.Sprintf("%s/%d", addr, product.ID)
		resp, err := http.Get(url)
		if err != nil {
			t.Fatalf(FailedToSendGetMessage, err)
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

func adminUpdateProduct(t *testing.T, product dto.ProductResponse, addr string) {
	t.Run("Update Product", func(t *testing.T) {
		updateRequest := dto.AdminUpdateProductRequest{
			Name:        "Updated Product",
			Description: "Updated description",
			Price:       1500,
			IsActive:    false,
		}

		body, err := json.Marshal(updateRequest)
		if err != nil {
			t.Fatalf("Failed to marshal update request: %v", err)
		}

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%d", addr, product.ID), bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("Failed to create PUT request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to send PUT request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
			responseBody, _ := io.ReadAll(resp.Body)
			t.Logf(ResponseBodyMessage, string(responseBody))
			return
		}

		var updatedProduct dto.ProductResponse
		if err := json.NewDecoder(resp.Body).Decode(&updatedProduct); err != nil {
			t.Fatalf("Failed to decode update response: %v", err)
		}

		if updatedProduct.Name != updateRequest.Name || updatedProduct.Description != updateRequest.Description ||
			updatedProduct.Price != updateRequest.Price {
			t.Errorf("Product update mismatch. Expected %+v, got %+v", updateRequest, updatedProduct)
		}
	})
}

func adminDeleteProduct(t *testing.T, product dto.ProductResponse, addr string) {
	t.Run("Delete Product", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%d", addr, product.ID), nil)
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
			t.Logf(ResponseBodyMessage, string(responseBody))
		}
	})
}
func adminVerifyProductCreated(t *testing.T, product dto.ProductResponse, addr string) {
	t.Run("Verify Product Deleted", func(t *testing.T) {
		resp, err := http.Get(fmt.Sprintf("%s/%d", addr, product.ID))
		if err != nil {
			t.Fatalf(FailedToSendGetMessage, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest && resp.StatusCode != http.StatusInternalServerError {
			t.Errorf("Expected error status after deletion, got %d", resp.StatusCode)
		}
	})
}
