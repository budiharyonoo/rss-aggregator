package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	fmt.Println("Test")

	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT is not bound in the env file")
	}

	fmt.Println("Port:", port)
}
