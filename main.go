package main

import (
	"./db"
	"./route"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init()
	route.Init()
}