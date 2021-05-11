package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log.Println("stating API cmd")
	port := os.Getenv("API_PORT")
	fmt.Println(port)
}
