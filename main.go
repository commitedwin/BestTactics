package main

import (
	"fmt"

	"github.com/commitedwin/BestTactics/riot"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	response := riot.GetChallengerLeague()
	for _, entry := range response.Entries {
		fmt.Println(entry.Puuid)
	}
}
