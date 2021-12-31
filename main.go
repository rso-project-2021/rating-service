package main

import (
	"log"
	"rating-service/config"
	"rating-service/db"
	"rating-service/server"
)

// @title rating-service API
// @version 1.0
// @description This is an API for ratings and comments.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http

func main() {

	// Load configuration settings.
	config, err := config.New(".")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Connect to the database.
	store, err := db.Connect(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Create a server and setup routes.
	server, err := server.NewServer(config, store)
	if err != nil {
		log.Fatal("Failed to create a server: ", err)
	}

	// Start a server.
	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("Failed to start a server: ", err)
	}
}
