package server

import (
	"go-reference-structure/internal/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
	DB   database.Database
}

func New(db database.Database) *Server {
	return &Server{
		Echo: echo.New(),
		DB:   db,
	}
}

func (server *Server) Start(addr string) error {

	srv := &http.Server{
		Addr:    addr,
		Handler: server.Echo,
	}

	return srv.ListenAndServe()
}
