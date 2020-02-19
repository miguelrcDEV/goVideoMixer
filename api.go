package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type streaming struct {
	Name string `json:"name"`
}

func health(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

func transcodeStreaming(w http.ResponseWriter, r *http.Request) {
	var newStreaming streaming
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
	}

	json.Unmarshal(reqBody, &newStreaming)
	log.Println(newStreaming)
	log.Println(newStreaming.Name)

	mediaPath := os.Getenv("MEDIA_PATH")
	inputPath := mediaPath + "/" + newStreaming.Name
	outputPath := mediaPath + "/" + newStreaming.Name + "/" + newStreaming.Name + ".mp4"

	transcode(inputPath, outputPath)
}
