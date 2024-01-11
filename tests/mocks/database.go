package mocks

import "context"

type MockDatabase struct{}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{}
}

// Make sure this mock implements the methods of the Database interface, found at internal/database/database.go

func (m *MockDatabase) Health(context.Context) error {
	return nil
}
