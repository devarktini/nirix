package web

import "github.com/gofiber/fiber/v3"

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
}

func (ws *WebServer) Start() error {
	return ws.server.Listen(":" + string(rune(ws.Port)))
}
