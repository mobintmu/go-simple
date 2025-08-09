package test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	application "go-simple/cmd/web/app"
)

func startTestServer() {
	a := application.New()
	a.Routes()
	a.StartServer()
	time.Sleep(200 * time.Millisecond) // Give server time to start
}

func TestHelloWorldEndpoint(t *testing.T) {
	t.Parallel()
	startTestServer()

	resp, err := http.Get("http://localhost:4000")
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

	// Parse JSON
	var data map[string]string
	err = json.Unmarshal(body, &data)
	if err != nil {
		t.Fatalf("Expected JSON response, got error: %v", err)
	}

	// Validate the message
	message, ok := data["message"]
	if !ok {
		t.Fatalf("Missing 'message' field in response: %v", data)
	}
	if message != "hello world" {
		t.Errorf("Expected message 'hello world', got '%s'", message)
	}
}
