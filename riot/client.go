package riot

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/commitedwin/BestTactics/config"
)

var API = config.LoadConfig

type LeagueEntry struct {
	Puuid string `json:"puuid"`
}

type ChallengerResponse struct {
	Entries []LeagueEntry `json:"entries"`
}

func GetChallengerLeague() ChallengerResponse {
	req, err := http.NewRequest("GET", "https://na1.api.riotgames.com/tft/league/v1/challenger", nil)
	if err != nil {
		fmt.Println("error creating request /tft/league/v1/challenger:", err)
	}
	req.Header.Set("X-Riot-Token", API().RiotAPIKey)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error making request /tft/league/v1/challenger", err)
	}
	defer res.Body.Close()

	challengerResponse := ChallengerResponse{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&challengerResponse); err != nil {
		fmt.Println("error decoding ChallengerResponse body")
	}
	return challengerResponse

}
