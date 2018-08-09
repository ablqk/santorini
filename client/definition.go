// Package client contains the functions to call the service.
package client

import (
	"github.com/ablqk/santorini/api"
)

// Client is an http client for the Santorini service.
type Client interface {
	api.Santorini
}

// New creates a new client.
func New(port int) (Client, error) {
	return client{}, nil
}
