package main

import (
	"basics_web/internal/apiserver"
	"basics_web/internal/config"
	"log"

	"github.com/BurntSushi/toml"
)

const (
	configFile = "./configs/blog.toml"
)

func main() {
	config := config.New()

	_, err := toml.DecodeFile(configFile, &config)
	if err != nil {
		log.Fatalln(err)
	}

	server, err := apiserver.New(config)
	if err != nil {
		log.Fatalln(err)
	}

	if err := server.Start(); err != nil {
		log.Fatalln(err)
	}
}
