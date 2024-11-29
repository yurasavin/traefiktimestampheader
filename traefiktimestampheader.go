// Package traefiktimestampheader is a plugin
package traefiktimestampheader

import (
	"context"
	"net/http"
	"time"
)

// Config holds the plugin configuration.
type Config struct{}

// CreateConfig creates and initializes the plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// RequestTimestamp is the struct implementing the Traefik plugin interface.
type RequestTimestamp struct {
	next http.Handler
}

// New creates a new RequestTimestamp plugin instance.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &RequestTimestamp{
		next: next,
	}, nil
}

// ServeHTTP adds the "request-received-at" header and forwards the request.
func (r *RequestTimestamp) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	timestamp := time.Now().UTC().Format(time.RFC3339Nano)
	req.Header.Set("x-request-received-at", timestamp)
	r.next.ServeHTTP(rw, req)
}
