package main

import (
	"Azgart/internal/interfaces/api"
	"log"
)

func main() {
	server := api.NewAPIServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
