package tests

import (
	"go-reference-structure/internal/handlers"
	"go-reference-structure/internal/server"
	"go-reference-structure/tests/helpers"
	"go-reference-structure/tests/mocks"
	"net/http"
	"testing"
)

// This test file isn't really that important, but the structure translates to testing the other endpoints too.

func TestHealth_HTTP(t *testing.T) {

	// Setup
	mockDb := mocks.NewMockDatabase()
	s := server.New(mockDb)

	healthHandler := handlers.NewHealth(s)

	c, _, rec := helpers.NewEchoContext(nil) // Request body is nil

	// Tests

	if err := healthHandler.Check(c); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", rec.Code)
	}

	if rec.Body.String() != "Healthy!" {
		t.Errorf("Expected body to be OK, got %v", rec.Body.String())
	}

}

func TestHealth_DB(t *testing.T) {

	// Setup
	mockDb := mocks.NewMockDatabase()
	s := server.New(mockDb)

	healthHandler := handlers.NewHealth(s)

	c, _, rec := helpers.NewEchoContext(nil) // Request body is nil

	// Tests

	if err := healthHandler.DB_Check(c); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", rec.Code)
	}

	if rec.Body.String() != "Healthy!" {
		t.Errorf("Expected body to be OK, got %v", rec.Body.String())
	}
}
