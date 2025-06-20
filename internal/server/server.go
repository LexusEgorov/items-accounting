package server

import "github.com/labstack/echo/v4"

type Server struct {
	server   echo.Echo
	handlers handlers
}

//TODO: start/stop

func New(productManager ProductManager, categoryManager CategoryManager) *Server {
	return &Server{
		server:   *echo.New(),
		handlers: *newHandlers(productManager, categoryManager),
	}
}
