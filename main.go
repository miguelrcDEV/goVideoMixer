package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

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

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", health)
	router.HandleFunc("/transcode", transcodeStreaming).Methods("POST")

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST).
	handler := cors.Default().Handler(router)

	fmt.Printf("Starting server on %v with stringPort %v\n", port, stringPort)

	http.ListenAndServe(stringPort, handler)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
