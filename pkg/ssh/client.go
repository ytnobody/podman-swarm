package ssh

import (
	"context"
)

type Client interface {
	Execute(ctx context.Context, cmd string) (string, error)
	Close() error
}

type ClientConfig struct {
	Host       string
	Port       int
	Username   string
	PrivateKey string
}

// NewClient creates a new SSH client
func NewClient(config ClientConfig) (Client, error) {
	// TODO: Implement SSH client initialization
	return nil, nil
}
