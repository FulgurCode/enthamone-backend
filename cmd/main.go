package main

import (
	"os"

	"github.com/FulgurCode/enthamone-backend/server"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load(".env")

	// Get port for the server
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "3000"
	}

	server.Run(port)
}
