package main

import (
	"./db"
	"./route"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	if err := godotenv.Load("default.env"); err != nil {
		log.Fatal("Error loading default.env file")
	}

	if _, err := os.Stat("./pic/"); os.IsNotExist(err) {
		os.Mkdir("./pic/", 0755)
	}

	db.Init()
	route.Init()
}