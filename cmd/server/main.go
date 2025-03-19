package main

import (
	_ "lego-api-go/docs"
	"lego-api-go/server"
	"log"
)

//	@title			Lego API
//	@version		1.0
//	@description	This is the documentation for Lego API REST API.
//
// @host	localhost:8000
func main() {
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
