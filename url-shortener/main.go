package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/udaichauhan/url_shortener/api/routes"
)

func main(){
	err := godotenv.Load();
	if err != nil {
		fmt.Println(err);
	}

	//we are using the GIN
	router := gin.Default();
	//now call here that steUpRouters function
	setUpRouters(router);
	
	port := os.Getenv("APP_PORT");
	if port == "" {
		port = "8080"
	}
	log.Fatal(router.Run(":"+port));
}

func setUpRouters(router *gin.Engine){
	router.POST("/api/v1", routes.ShortenURL);
	router.GET("/api/v1/:shortID", routes.GetByShortID);
}