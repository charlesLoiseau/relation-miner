package main

import (
	"fmt"
	"log"
	"os"
	"relation-miner/neo4j"
	"relation-miner/twitter"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	arg := os.Args
	if len(arg) != 2 {
		fmt.Println("USAGE: ", os.Args[0], " user")
		return
	}
	n := neo4j.Neo4j{}
	t := twitter.Twitter{AccountName: os.Args[1]}
	t.GetRelation()
	println("Follower: ", len(t.Follower), " Following: ", len(t.Following), "target: ", t.User.Id, t.User.Name)
	n.User = t.User
	n.Following = t.Following
	n.Follower = t.Follower
	n.Init()
	n.CreateRelation()
	n.Close()
}
