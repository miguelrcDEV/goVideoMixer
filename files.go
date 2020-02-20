package main

import (
	"fmt"
	"io"
	"os"
	"log"
)

func createFile(path string) {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) { return }
		defer file.Close()
	}

	fmt.Println("==> done creating file", path)
}

func writeFile(path string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) { return }
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString("halo\n")
	if isError(err) { return }
	_, err = file.WriteString("mari belajar golang\n")
	if isError(err) { return }

	// save changes
	err = file.Sync()
	if isError(err) { return }

	fmt.Println("==> done writing to file")
}

func readFile(path string) {
	// re-open file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) { return }
	defer file.Close()

	// read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)
		
		// break if finally arrived at end of file
		if err == io.EOF {
			break
		}
		
		// break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}
	
	fmt.Println("==> done reading from file")
	fmt.Println(string(text))
}

func deleteFile(path string) {
	// delete file
	var err = os.Remove(path)
	if isError(err) { return }

	fmt.Println("==> done deleting file")
}

func appendBytes(path string, bytes []byte){
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	if _, err := f.Write(bytes); err != nil {
		log.Println(err)
	}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}