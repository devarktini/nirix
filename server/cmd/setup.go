package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/devarktini/nirix/server/common"
	"github.com/devarktini/nirix/server/internal/db"
	"github.com/devarktini/nirix/server/internal/web"
)

func setup() {
	printBanner()
	common.GetLogger().Log.Info("Starting the setup of NIRIX Server")
	db.Setup()
	db.GetInstance().Migrate()
	ws := web.NewWebServer()
	ws.Setup()
	sigCh := make(chan os.Signal, 1)
	go func() {
		if err := ws.Start(); err != nil {
			common.GetLogger().Log.Sugar().Fatalf("failed to start web server: %s", err.Error())
		}
	}()

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	log.Println("Shutting down server...")

	if err := ws.ShutdownWithTimeout(30 * time.Second); err != nil {
		ws.Shutdown()
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
