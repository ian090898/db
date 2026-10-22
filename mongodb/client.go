package mongodb

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/free5gc/db"
	"github.com/free5gc/util/mongoapi"
)

// Phase 1 uses the existing mongoapi global client.
// These variables manage its lifetime within the db package.
var (
	connectionMu   sync.Mutex
	referenceCount int
	currentDBName  string
	currentURL     string
)

// Client represents one user of the shared MongoDB connection.
type Client struct {
	closed bool
}

// Connect initializes or reuses the existing mongoapi connection.
func Connect(ctx context.Context, cfg *db.MongodbConfig) (*Client, error) {
	if ctx == nil {
		return nil, errors.New("mongodb: nil context")
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if cfg == nil || cfg.Name == "" || cfg.Url == "" {
		return nil, errors.New("mongodb: name and url are required")
	}

	connectionMu.Lock()
	defer connectionMu.Unlock()

	if referenceCount > 0 {
		if currentDBName != cfg.Name || currentURL != cfg.Url {
			return nil, errors.New("mongodb: conflicting connection configuration")
		}
	} else if mongoapi.Client != nil {
		return nil, errors.New("mongodb: mongoapi client was initialized outside db package")
	}

	if err := mongoapi.SetMongoDB(cfg.Name, cfg.Url); err != nil {
		return nil, fmt.Errorf("mongodb: connect: %w", err)
	}

	if referenceCount == 0 {
		currentDBName = cfg.Name
		currentURL = cfg.Url
	}
	referenceCount++

	return &Client{}, nil
}

// Close releases this user's reference to the shared connection.
// The actual MongoDB connection closes when its last user closes.
func (c *Client) Close(ctx context.Context) error {
	if c == nil {
		return nil
	}
	if ctx == nil {
		return errors.New("mongodb: nil context")
	}

	connectionMu.Lock()
	defer connectionMu.Unlock()

	if c.closed {
		return nil
	}

	if referenceCount == 1 && mongoapi.Client != nil {
		if err := mongoapi.Client.Disconnect(ctx); err != nil {
			return fmt.Errorf("mongodb: disconnect: %w", err)
		}
		mongoapi.Client = nil
	}

	if referenceCount > 0 {
		referenceCount--
	}
	if referenceCount == 0 {
		currentDBName = ""
		currentURL = ""
	}

	c.closed = true
	return nil
}
