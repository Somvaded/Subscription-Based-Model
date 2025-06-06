package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Somvaded/subscription-management/config"
	"github.com/Somvaded/subscription-management/database"
	"github.com/Somvaded/subscription-management/repository"
	"github.com/Somvaded/subscription-management/routes"
	"github.com/gin-gonic/gin"
)
func main(){
	// Load configuration
	cfg , err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	// Connect to the database with retry logic
	db, err := database.ConnectToDBwithRetry(cfg.DatabaseURL, 5)
	if err != nil {
		log.Fatalf("Failed to connect to database after retries: %v", err)
	}
	// Set up Gin router and routes

	r := gin.Default()
	routes.SetupRoutes(r,db)

	// Launch expiry job
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		for range ticker.C {
			repository.ExpireDueSubscriptions(db)
		}
	}()
	
	// Start the server
	log.Fatal(r.Run(fmt.Sprintf(":%s", cfg.Port)))
}