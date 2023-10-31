package main

import (
	"log"
)

var (
	GET    = "GET"
	POST   = "POST"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":8081", store)
	server.Run()
}
