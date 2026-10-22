package mongodb

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/free5gc/db"
	"github.com/free5gc/util/mongoapi"
)

// TestMongoDBConnectionLifecycle exercises the Phase 1 shared mongoapi client.
// Set MONGODB_URL to a running MongoDB URI. It is not run in parallel.
func TestMongoDBConnectionLifecycle(t *testing.T) {
	url := os.Getenv("MONGODB_URL")
	if url == "" {
		t.Skip("set MONGODB_URL to run the MongoDB integration test")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cfg := &db.MongodbConfig{
		Name: "free5gc_db_m0_test",
		Url:  url,
	}

	first, err := Connect(ctx, cfg)
	if err != nil {
		t.Fatalf("first Connect: %v", err)
	}
	t.Cleanup(func() { _ = first.Close(context.Background()) })

	if mongoapi.Client == nil {
		t.Fatal("Connect returned a nil mongoapi.Client")
	}
	original := mongoapi.Client
	if err := original.Ping(ctx, nil); err != nil {
		t.Fatalf("MongoDB Ping after first Connect: %v", err)
	}

	second, err := Connect(ctx, cfg)
	if err != nil {
		t.Fatalf("second Connect: %v", err)
	}
	t.Cleanup(func() { _ = second.Close(context.Background()) })
	if mongoapi.Client != original {
		t.Fatal("second Connect unexpectedly replaced the shared client")
	}

	// A conflicting configuration must not replace the shared connection.
	_, err = Connect(ctx, &db.MongodbConfig{Name: "different_db", Url: url})
	if err == nil {
		t.Fatal("expected error for conflicting MongoDB configuration")
	}

	if err := first.Close(ctx); err != nil {
		t.Fatalf("first Close: %v", err)
	}
	if mongoapi.Client != original {
		t.Fatal("first Close must keep the shared client alive")
	}
	if err := mongoapi.Client.Ping(ctx, nil); err != nil {
		t.Fatalf("MongoDB Ping after first Close: %v", err)
	}

	if err := second.Close(ctx); err != nil {
		t.Fatalf("second Close: %v", err)
	}
	if mongoapi.Client != nil {
		t.Fatal("last Close must release mongoapi.Client")
	}
	if err := second.Close(ctx); err != nil {
		t.Fatalf("repeated Close must be idempotent: %v", err)
	}

	// A new connection must work after the last user has closed the old one.
	third, err := Connect(ctx, cfg)
	if err != nil {
		t.Fatalf("Connect after Close: %v", err)
	}
	t.Cleanup(func() { _ = third.Close(context.Background()) })
	if mongoapi.Client == nil {
		t.Fatal("Connect after Close returned a nil mongoapi.Client")
	}
	if err := mongoapi.Client.Ping(ctx, nil); err != nil {
		t.Fatalf("MongoDB Ping after reconnect: %v", err)
	}
	if err := third.Close(ctx); err != nil {
		t.Fatalf("final Close: %v", err)
	}
}
