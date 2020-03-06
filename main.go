package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Server
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	fmt.Println("Hoş geldin")
}
