package main

import (
	"basics_web/internal/apiserver"
	"log"
)

func main() {
	server := apiserver.New()
	if err := server.Start(); err != nil {
		log.Fatalln(err)
	}
}
