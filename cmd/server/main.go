package main

import (
	"log"

	"lego-api-go/server"
)

func main() {
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
