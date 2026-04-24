package config

import (
	"log"
	"os"
)

type Config struct {
	RiotAPIKey string
	DBUrl      string
}

func LoadConfig() Config {
	riotAPIKey, exists := os.LookupEnv("RIOT_API_KEY")

	if !exists {
		log.Fatal("RIOT_API_KEY not found")
	}
	dbURL, exists := os.LookupEnv("DB_URL")
	if !exists {
		log.Fatal("DB_URL not found")
	}
	return Config{
		RiotAPIKey: riotAPIKey,
		DBUrl:      dbURL,
	}

}
