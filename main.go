package main

import (
	"fmt"
	"log"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func printName() {
	myFigure := figure.NewFigure("MERGETS", "", true)
	myFigure.Print()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file loaded")
	}

	// configure the hls path and server port
	port := os.Getenv("PORT")
	stringPort := fmt.Sprintf(":%v", port)
	mediaPath := os.Getenv("MEDIA_PATH")
	fmt.Printf("Media path on %v\n", mediaPath)

	// SETUP API
	router := gin.Default()
	router.Use(gin.Recovery()) // Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(cors.Default())

	api := router.Group("/api")
	{
		api.GET("/health", health)
		api.POST("/transcode", transcodeStreaming)
	}

	printName()
	log.Printf("Media path is %s\n", mediaPath)

	router.Run(stringPort)
}
