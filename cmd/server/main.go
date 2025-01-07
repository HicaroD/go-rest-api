package main

import (
	"log"

	"github.com/HicaroD/api/server"
)

func main() {
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
