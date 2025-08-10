package test

import (
	"net/http"
	"testing"
)

func TestSlowEndpoint(t *testing.T) {
	t.Parallel()
	startTestServer()

	resp, err := http.Get("http://localhost:4000/api/v1/slow")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusRequestTimeout {
		t.Errorf("Expected status 408 OK, got %d", resp.StatusCode)
	}
}
