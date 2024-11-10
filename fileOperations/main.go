package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Constants
const sagar = "sagar.txt"
const url = "https://raw.githubusercontent.com/sagar0419/sagar0419/refs/heads/master/README.md"
const gitFile = "README.md"

// Creating a file
func createFile(s string) {
	fmt.Println("Create Operation")
	fmt.Println("Creating a text file")

	// Creating a file
	file, err := os.Create(s)
	if err != nil {
		log.Fatal("error occured : ", err)
		return
	}

	defer file.Close()

	// Writing a file
	_, err = file.WriteString("This file is created using golang code")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("File Name is: %s\n\n", file.Name())
}

// Reading a existing file
func readFile(s string) {
	fmt.Println("Read Operation")
	fmt.Println("Reading the previously created file")

	// Reading a file
	content, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("content of the file is: -  %s\n\n", content)
}

// Appending a existing file
func appendFile(s string) {
	fmt.Println("Updating file")
	fmt.Println("Updating the previously created file")

	// opening file to append
	file, err := os.OpenFile(s, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// adding new line
	_, err = file.WriteString("\nadded a new line")
	if err != nil {
		log.Fatal(err)
		return
	}

	newContent, err := os.ReadFile(s)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Updated file content: %s\n\n", string(newContent))
}

// Read a remote file

func remoteRead(s string) {
	fmt.Printf("Reading this remote url %s", url)

	// accessing a URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Error: Received status code %d from URL", resp.StatusCode)
	}
	// Reading Response
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Content of the file is \n\n %s\n\n this content is also written in the file named %s\n", string(content), gitFile)

	// Creating a file and writing the upstream content to it
	file, err := os.Create(s)
	if err != nil {
		log.Fatal("error occured : ", err)
		return
	}

	defer file.Close()
	_, err = file.WriteString(string(content))
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	createFile(sagar)
	readFile(sagar)
	appendFile(sagar)
	remoteRead(gitFile)
}
