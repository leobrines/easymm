package steam

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetPlayerSummaries(steamId string) (*PlayerSummaries, error) {
	url := fmt.Sprintf(steamAPIGetPlayerSummariesURL, apiKey, steamId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type Result struct {
		Response struct {
			Players []PlayerSummaries `json:"players"`
		} `json:"response"`
	}
	var data Result
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return &data.Response.Players[0], err
}
