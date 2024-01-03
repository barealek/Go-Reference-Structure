package routes

import (
	"go-reference-structure/internal/handlers"
	s "go-reference-structure/internal/server"
)

func Register(server *s.Server) {
	RegisterMiddlewares(server)

	authHandler := handlers.NewAuth(server)
	healthHandler := handlers.NewHealth(server)

	auth := server.Echo.Group("/auth")
	auth.GET("/login", authHandler.Login)
	auth.GET("/logout", authHandler.Logout)

	health := server.Echo.Group("/health")
	health.GET("/", healthHandler.Check)
	health.GET("/db", healthHandler.DB_Check)

}
