package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"payment/internal/config"
	"syscall"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "conf", "configs/config.toml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	c, err := config.New(confPath)
	if err != nil {
		panic(fmt.Sprintf("error configuring service: %v", err))
	}

	app, cleanup, err := initApp(c.Logger, c.Server, c.Data)
	if err != nil {
		panic(fmt.Sprintf("error initializing app: %v", err))
	}
	defer func() {
		cleanup()
		app.Shutdown()
	}()
	errChan := make(chan error)
	go func() {
		if err := app.ServeGRPC(); err != nil {
			errChan <- fmt.Errorf("serving gRPC: %w", err)
		}
	}()
	go func() {
		if err := app.ServeHTTP(); err != nil {
			errChan <- fmt.Errorf("serving HTTP: %w", err)
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	select {
	case err := <-errChan:
		panic(err)
	case <-stopChan:
	}
}
