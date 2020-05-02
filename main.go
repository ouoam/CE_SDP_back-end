package main

import (
	"./db"
	"./route"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	godotenv.Load()
	if err := godotenv.Load("default.env"); err != nil {
		log.Fatal("Error loading default.env file")
	}

	db.Init()
	route.Init()
}