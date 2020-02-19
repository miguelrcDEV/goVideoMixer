package main

import (
	"fmt"
	"github.com/xfrr/goffmpeg/transcoder"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	tsExt           = ".ts"
	mergeTSFilename = "merged"
	progressWidth   = 40
)

func joinTsFiles(inputPath string) []byte {
	var allTs []byte

	totalTsFiles := 0
	filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), tsExt) {
			totalTsFiles++
		}
		return nil
	})

	log.Println("TOTAL TS FILES TO CONCATENATE => " + strconv.Itoa(totalTsFiles))

	mergedCount := 0
	filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".ts") {
			println("concatenating " + path + " ...")
			b, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			allTs = append(allTs, b...)

			mergedCount++
			DrawProgressBar("Merge .ts files", float32(mergedCount)/float32(totalTsFiles), progressWidth)
		}
		return nil
	})
	return allTs
}

func transcode(inputPath string, outputPath string) {
	allTsPath := fmt.Sprintf("%s/%s", inputPath, mergeTSFilename)

	allTs := joinTsFiles(inputPath)
	log.Println("ALL TS FILE PATH" + allTsPath)
	ioutil.WriteFile(allTsPath, allTs, 0644)

	// Create new instance of transcoder
	trans := new(transcoder.Transcoder)

	// Initialize transcoder passing the input file path and output file path
	err := trans.Initialize(allTsPath, outputPath)
	// Handle error...
	if err != nil {
		log.Fatal(err)
	}

	// Start transcoder process with progress checking
	done := trans.Run(true)

	// Returns a channel to get the transcoding progress
	progress := trans.Output()

	// Example of printing transcoding progress
	for progressValues := range progress {
		/*log.Println(progressValues.Progress)
		log.Println(float32(progressValues.Progress))*/
		DrawProgressBar("Merge .ts files", float32(progressValues.Progress)/100, progressWidth)
		/*log.Println("FRAMES " + progressValues.FramesProcessed)
		log.Println("CURRENT TIME " + progressValues.CurrentTime)
		log.Println("CURRENT BITRATE " + progressValues.CurrentBitrate)
		log.Println("SPEED " + progressValues.Speed)*/
	}

	log.Println("TRANSCODING DONE")

	// This channel is used to wait for the transcoding process to end
	err = <-done
}
