package main

import (
	"log"
	"pajarit-feed-service/config"
	"pajarit-feed-service/server"
)

func main() {
	println("loading configuration")

	cfg, err := config.LoadConfiguration()
	if err != nil {
		log.Fatalln("can't load configuration")
	}

	println("loading dependencies")

	deps, err := config.BuildDependencies(cfg)
	if err != nil {
		log.Fatalln("can't load dependencies")
	}

	err = server.StartServer(cfg, deps)
	if err != nil {
		log.Fatalln("can't start server")
	}
}
