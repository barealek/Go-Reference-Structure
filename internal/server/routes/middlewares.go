package routes

import (
	s "go-reference-structure/internal/server"

	"github.com/barealek/echowares/echologger"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterMiddlewares(server *s.Server) {
	// This is where you would register app-wide middlewares,
	// such as CORS, logging, etc.

	server.Echo.Use(middleware.AddTrailingSlash())

	server.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "https://production-site.com"},
	}))

	server.Echo.Use(echologger.New())
}
