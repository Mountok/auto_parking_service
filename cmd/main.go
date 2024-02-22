package main

import (
	"context"
	"go_service_parking/example/internals/app"
	"go_service_parking/example/internals/cfg"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {
	config := cfg.LoadAndStoreConfig()

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	server := app.NewServer(config, ctx)

	go func() {
		oscall := <-c
		log.Printf("Sytem call: %+v",oscall)
		server.Shutdown()
		cancel()
	}()

	server.Serve()
}
