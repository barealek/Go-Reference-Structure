package database

import (
	"context"
)

type Database interface {
	Health(context.Context) error
	// UpdateUser(models.User) error
	// CheckAPIKey(apiKey string) bool
	// GetUser(userid string) (models.User, error)

	// The database interface will implement the methods that are commonly used by the application.
}
