package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/anton/config"
	"github.com/vishalrana9915/anton/db"
)

func Execute() {

	// Load configuration
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database connection
	dbConn, err := db.NewConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
	// Example of starting your application (e.g., API server or worker)
	log.Println("Application is running...")

}
