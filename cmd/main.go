package main

import (
	"htn/server"

	dotenv "github.com/joho/godotenv"
)

func main() {
	// Read .env
	dotenv.Load(".env")
	server.Setup()
}
