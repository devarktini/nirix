package web

import (
	"github.com/devarktini/nirix/server/internal/routes"
	"github.com/gofiber/fiber/v3"
)

type WebServer struct {
	Port   int
	server *fiber.App
}

func NewWebServer(port int) *WebServer {
	return &WebServer{
		Port:   port,
		server: fiber.New(fiber.Config{}),
	}
}

func (ws *WebServer) Setup() {
	// Setup routes and middleware here

	v1Group := ws.server.Group("/v1")
	routes.SetupAuthRoutes(v1Group)
}

func (ws *WebServer) Start() error {
	return ws.server.Listen(":" + string(rune(ws.Port)))
}
