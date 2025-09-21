package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/udaichauhan/url_shortener/api/database"
	"github.com/udaichauhan/url_shortener/api/routes"
)

func main(){
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found:", err)
	}

	// Set Gin to debug mode for better error visibility
	gin.SetMode(gin.DebugMode);

	database.InitializeClient();

	// Create router with default middleware (Logger and Recovery)
	router := gin.Default()
	
	// Add custom recovery middleware for better error handling
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(500, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(500)
	}))

	// Setup routes
	setUpRouters(router)
	
	// Get port from environment
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	
	fmt.Printf("Server starting on port %s...\n", port)
	fmt.Printf("Environment variables:\n")
	fmt.Printf("- APP_PORT: %s\n", os.Getenv("APP_PORT"))
	fmt.Printf("- API_QUOTA: %s\n", os.Getenv("API_QUOTA"))
	fmt.Printf("- DOMAIN: %s\n", os.Getenv("DOMAIN"))
	
	log.Fatal(router.Run(":" + port))
}

func setUpRouters(router *gin.Engine){
	// Add a health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	
	router.POST("/api/v1", routes.ShortenURL)
	router.GET("/api/v1/:shortID", routes.GetByShortID)
	router.DELETE("/api/v1/:shortID", routes.DeleteURL)
	router.PUT("/api/v1/:shortID", routes.EditURL)
	router.POST("/api/v1/addTag", routes.AddTag)
}