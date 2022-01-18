package main

import (
	"flag"
	"fmt"
	"payment/internal/config"
)

var (
	confPath string
)

func init() {
	flag.StringVar(&confPath, "conf", "configs/config.toml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	c, err := config.Configure(confPath)
	if err != nil {
		panic(fmt.Sprintf("error configuring service: %v", err))
	}

	app, cleanup, err := initApp(c.Logger, c.Server, c.Data)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	if err := app.Serve(); err != nil {
		panic(err)
	}
}
