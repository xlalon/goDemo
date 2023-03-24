package main

import (
	"flag"
	"log"

	"github.com/xlalon/golee/internal/domain/wallet/conf"
	"github.com/xlalon/golee/internal/domain/wallet/http"
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
