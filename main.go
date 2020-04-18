package main

import (
	"log"
	"relation-miner/twitter"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	t := twitter.Twitter{
		AccountName: "cha93100",
	}
	t.GetRelation()

}
