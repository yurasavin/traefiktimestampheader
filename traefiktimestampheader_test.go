package traefiktimestampheader

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTimestampHeader(t *testing.T) {
	cfg := CreateConfig()
	ctx := context.Background()

	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	// Create the plugin
	plugin, err := New(ctx, next, cfg, "timestamp-header")
	if err != nil {
		t.Fatalf("Failed to create plugin: %v", err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the plugin
	plugin.ServeHTTP(recorder, req)

	// Validate the header
	headerValue := req.Header.Get("x-request-received-at")
	if headerValue == "" {
		t.Fatalf("Expected 'request-received-at' header, but got none")
	}

	// Parse and validate the timestamp format
	_, err = time.Parse(time.RFC3339Nano, headerValue)
	if err != nil {
		t.Errorf("Invalid timestamp format in header: %v", err)
	}
}
