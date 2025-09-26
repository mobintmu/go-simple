package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	app "go-simple/internal/app"
	"go-simple/internal/config"

	"go.uber.org/fx"
)

func StartHTTPServer() *fx.App {
	a := app.NewApp()
	go a.Run()
	time.Sleep(300 * time.Millisecond) // give server time to start
	return a
}

func WithTestServer(t *testing.T, testFunc func()) {
	a := StartHTTPServer()
	defer a.Stop(context.Background())
	testFunc()
}

func TestHealthEndpoint(t *testing.T) {
	t.Parallel()
	WithTestServer(t, func() {
		cfg, err := config.NewConfig()
		if err != nil {
			t.Fatalf("Failed to load config: %v", err)
		}
		addr := fmt.Sprintf("http://%s:%d", cfg.HTTPAddress, cfg.HTTPPort)
		resp, err := http.Get(addr + "/health")
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
		if message != "OK" {
			t.Errorf("Expected message 'hello world', got '%s'", message)
		}
	})
}
