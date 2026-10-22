// Package dbif defines database-independent operations for free5GC NFs.
//
// Interfaces must not expose MongoDB driver types, collection names,
// or HTTP ProblemDetails. Every operation takes context.Context first.
package dbif
