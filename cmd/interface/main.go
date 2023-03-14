package main

import (
	"flag"
	"log"

	"github.com/xlalon/golee/internal/interface/conf"
	"github.com/xlalon/golee/internal/interface/http"
	"github.com/xlalon/golee/pkg/net/http/server"
)

func main() {

	flag.Parse()

	if err := conf.Init(); err != nil {
		log.Fatal(err)
	}

	engine := server.NewEngine(conf.Conf.Server)

	if err := http.Init(engine, conf.Conf); err != nil {
		log.Fatal(err)
	}

	engine.RunServer(conf.Conf.Server)
}
