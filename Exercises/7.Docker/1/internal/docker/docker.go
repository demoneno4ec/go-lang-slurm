package docker

import (
	"context"

	"github.com/docker/docker/client"
)

const (
	capStatCh = 10
)

// Client wrapper around docker sdk client.
type Client struct {
	c *client.Client
}

// NewDockerClient builder for DockerClient.
func New() (*Client, error) {
	// TODO: код писать здесь
}

type Container struct {
	ID   string
	Name string
}

// ListCID returns slice running docker containers ids.
func (c *Client) ListCID(ctx context.Context) ([]Container, error) {
	// TODO: код писать здесь
}
