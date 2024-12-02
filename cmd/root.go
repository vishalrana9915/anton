package cmd

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/anton/config"
	"github.com/vishalrana9915/anton/db"
)

func Execute() {

	// Load configuration
	config.LoadConfig()

	fmt.Printf("Starting up ->%s \n", config.GetConfig("APP_NAME", ""))

	// Initialize database connection
	dbConn, err := db.NewConnection(config.GetConfig("DATABASE_URL", ""))
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

	r.Run(config.GetConfig("APP_PORT", ":3005")) // listen and serve on
	// Example of starting your application (e.g., API server or worker)
	log.Println("Application is running...")

}
