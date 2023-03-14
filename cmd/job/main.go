package main

import (
	"flag"
	"log"

	"github.com/xlalon/golee/internal/job"
	"github.com/xlalon/golee/internal/job/conf"
	"github.com/xlalon/golee/pkg/job/worker"
)

func main() {

	flag.Parse()

	if err := conf.Init(); err != nil {
		log.Fatal(err)
	}

	server := worker.NewServer(conf.Conf.Job)

	if err := job.Init(server, conf.Conf); err != nil {
		log.Fatal(err)
	}

	if err := server.RunWorker(conf.Conf.Job); err != nil {
		log.Fatal(err)
	}

}
