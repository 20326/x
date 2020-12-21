package redis

import (
	"context"
	"sync"

	r "github.com/go-redis/redis/v8"
)

type Client struct {
	Cli r.UniversalClient
	Ctx context.Context
	m   sync.RWMutex
}

// NewUniversalClient create a new Redis cache instance.
func NewUniversalClient(ctx context.Context, addrs []string, opts ...OptionsFunc) (*Client, error) {
	client := &Client{
		Ctx: ctx,
	}

	if client.Ctx == nil {
		client.Ctx = context.Background()
	}

	client.m.Lock()
	defer client.m.Unlock()

	uo := &r.UniversalOptions{
		Addrs:         addrs,
		RouteRandomly: true,
	}

	for _, opt := range opts {
		opt(uo)
	}

	client.Cli = r.NewUniversalClient(uo)

	return client, nil
}
