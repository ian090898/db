package mongodb

import (
	"context"
	"testing"

	"github.com/free5gc/db"
)

func TestConnectRejectsInvalidConfig(t *testing.T) {
	tests := []struct {
		name string
		cfg  *db.MongodbConfig
	}{
		{
			name: "nil config",
			cfg:  nil,
		},
		{
			name: "missing database name",
			cfg: &db.MongodbConfig{
				Url: "mongodb://127.0.0.1:27017",
			},
		},
		{
			name: "missing URL",
			cfg: &db.MongodbConfig{
				Name: "free5gc",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Connect(context.Background(), tt.cfg)
			if err == nil {
				t.Fatal("expected an error for invalid config")
			}
		})
	}
}
