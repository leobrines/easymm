package steam

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/leobrines/easymm/pkg"
	"github.com/leobrines/easymm/pkg/util"
)

type Service struct {
	apiKey string
}

func NewService(apiKey string) *Service {
	if apiKey == "" {
		log.Fatal("Steam API Key must be non-empty")
	}
	return &Service{apiKey}
}

func (s *Service) UserLogin(r *http.Request) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	openID := NewOpenId(r)

	switch openID.Mode() {
	case "":
		result["steam_login_url"] = openID.AuthUrl()
	default:
		steamID, err := openID.ValidateAuth()
		if err != nil {
			return nil, err
		}

		player, err := s.getPlayerSummaries(steamID)
		if err != nil {
			return nil, err
		}

		result = util.StructToMap(player)
	}

	return result, nil
}

func (s *Service) getPlayerSummaries(steamId string) (*PlayerSummaries, error) {
	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", s.apiKey, steamId)
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
