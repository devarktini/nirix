package web

import (
	"strconv"
	"time"

	"github.com/devarktini/nirix/server/common"
	"github.com/devarktini/nirix/server/config"
	"github.com/devarktini/nirix/server/internal/routes"
	"github.com/gofiber/fiber/v3"
)

type WebServer struct {
	Port   int
	server *fiber.App
}

func NewWebServer() *WebServer {
	return &WebServer{
		Port:   config.GetConfig().ServerPort,
		server: fiber.New(fiber.Config{}),
	}
}

func (ws *WebServer) Setup() {
	// Setup routes and middleware here

	v1Group := ws.server.Group("/v1")
	routes.SetupAuthRoutes(v1Group)
	routes.SetupUserRoutes(v1Group)
	routes.SetupOpsRoutes(v1Group)
}

func (ws *WebServer) Start() error {
	common.GetLogger().Log.Sugar().Infof("Starting web server on port %d\n", ws.Port)
	return ws.server.Listen(":" + strconv.Itoa(ws.Port))
}

func (ws *WebServer) ShutdownWithTimeout(timeOut time.Duration) error {
	return ws.server.ShutdownWithTimeout(timeOut)
}

func (ws *WebServer) Shutdown() error {
	return ws.server.Shutdown()
}
