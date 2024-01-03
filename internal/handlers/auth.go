package handlers

import (
	s "go-reference-structure/internal/server"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	server *s.Server
}

func NewAuth(server *s.Server) *AuthHandler {
	return &AuthHandler{
		server: server,
	}
}

func (s *AuthHandler) Login(c echo.Context) error {
	return c.String(200, "Login!")
}

func (s *AuthHandler) Logout(c echo.Context) error {
	return c.String(200, "Logout!")
}
