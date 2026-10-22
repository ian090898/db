package dbfactory

import (
	"testing"

	"github.com/free5gc/db"
)

func TestValidateConfig(t *testing.T) {
	valid := &db.Config{
		Type: db.DbTypeMongoDB,
		Mongodb: &db.MongodbConfig{
			Name: "free5gc",
			Url:  "mongodb://127.0.0.1:27017",
		},
	}

	tests := []struct {
		name    string
		cfg     *db.Config
		wantErr bool
	}{
		{"nil config", nil, true},
		{"unsupported type", &db.Config{Type: "postgresql"}, true},
		{"missing mongodb config", &db.Config{Type: db.DbTypeMongoDB}, true},
		{
			"missing name",
			&db.Config{
				Type: db.DbTypeMongoDB,
				Mongodb: &db.MongodbConfig{
					Url: "mongodb://127.0.0.1:27017",
				},
			},
			true,
		},
		{
			"missing URL",
			&db.Config{
				Type:    db.DbTypeMongoDB,
				Mongodb: &db.MongodbConfig{Name: "free5gc"},
			},
			true,
		},
		{"valid mongodb", valid, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateConfig(tt.cfg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("validateConfig() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
