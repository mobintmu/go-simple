package test

import (
	"io"
	"net/http"
	"testing"
	"time"

	application "go-simple/cmd/web/app"
)

func startTestServer() {
	a := application.New()
	a.StartServer()
	time.Sleep(200 * time.Millisecond) // Give server time to start
}

func TestHelloWorldEndpoint(t *testing.T) {
	t.Parallel()
	startTestServer()

	resp, err := http.Get("http://localhost:4000/")
	if err != nil {
		t.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	expected := "Hello World"
	if string(body) != expected {
		t.Errorf("Expected body '%s', got '%s'", expected, string(body))
	}
}
