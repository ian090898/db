package dbfactory

import (
	"errors"
	"fmt"

	"github.com/free5gc/db"
)

// validateConfig validates the common database configuration.
// NF-specific constructors will be added as each NF is migrated.
func validateConfig(cfg *db.Config) error {
	if cfg == nil {
		return errors.New("dbfactory: nil database config")
	}

	switch cfg.Type {
	case db.DbTypeMongoDB:
		if cfg.Mongodb == nil {
			return errors.New("dbfactory: missing mongodb config")
		}
		if cfg.Mongodb.Name == "" || cfg.Mongodb.Url == "" {
			return errors.New("dbfactory: mongodb name and url are required")
		}
		return nil

	default:
		return fmt.Errorf(
			"dbfactory: unsupported database type %q",
			cfg.Type,
		)
	}
}
