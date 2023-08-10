package database

import "github.com/codeinuit/fizzbuzz-api/pkg/models"

// Database represent the interface implementation
// for database operations
type Database interface {
	CountUsage() (models.Stats, error)
	UsageUpdate(m models.Stats)
}
