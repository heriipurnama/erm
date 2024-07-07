package main

import (
	"dbo/erm/config"
	"dbo/erm/routes"
	"log"
)

func main() {
	// Connect to the database
	_, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		panic(err)
	}

	// Setup router
	r := routes.SetupRouter()

	// Run the server
	r.Run(":8080")
}
