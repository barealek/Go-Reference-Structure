package handlers

import (
	"net/http"

	s "go-reference-structure/internal/server"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	server *s.Server
}

// Each handler should have a New function that takes the server as an argument
// and returns a pointer to the handler.

// This allows us to inject dependencies into the handler as needed,
// and makes it easier to test.

func NewHealth(server *s.Server) *HealthHandler {
	return &HealthHandler{
		server: server,
	}
}

func (s *HealthHandler) Check(c echo.Context) error {
	return c.String(http.StatusOK, "Healthy!")
}

func (s *HealthHandler) DB_Check(c echo.Context) error {

	if err := s.server.DB.Health(c.Request().Context()); err != nil {
		return c.String(http.StatusInternalServerError, "Unhealthy!")
	}

	return c.String(http.StatusOK, "Healthy!")
}
