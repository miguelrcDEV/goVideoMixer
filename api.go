package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type streaming struct {
	Name string `json:"name"`
}

func health(c *gin.Context) {
	c.JSON(200, gin.H{
		"alive": "true",
	})
}

func transcodeStreaming(context *gin.Context) {
	var newStreaming streaming
	name := context.PostForm("name")
	newStreaming.Name = name

	mediaPath := os.Getenv("MEDIA_PATH")
	inputPath := mediaPath + "/" + newStreaming.Name
	outputPath := mediaPath + "/" + newStreaming.Name + "/" + newStreaming.Name + ".mp4"

	log.Printf("INPUT PATH => %v", inputPath)
	log.Printf("OUTPUT PATH => %v", outputPath)

	existsInputDir := ExistsDir(inputPath)

	if existsInputDir {
		transcode(inputPath, outputPath)
	} else {
		context.JSON(404, gin.H{
			"NOT_FOUND": "Dir for " + name,
		})
	}
}
