package db

import "errors"

// Document represents a schemaless database record.
type Document = map[string]interface{}

var (
	ErrNotFound = errors.New("db: not found")

	ErrAlreadyExists = errors.New("db: already exists")
)
