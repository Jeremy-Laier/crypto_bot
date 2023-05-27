package discord

import (
	"bytes"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

type Command struct {
	Name        string `json:"name"`
	Type        int    `json:"type"`
	Description string `json:"description"`
	Options     []struct {
		Name        string      `json:"name"`
		Description string      `json:"description"`
		Type        int         `json:"type"`
		Required    interface{} `json:"required"`
		Choices     []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"choices"`
	} `json:"options"`
}

func (d discordImpl) AddCommand(command Command) error {
	route := "commands"

	url := fmt.Sprintf("%v/%v/%v", d.URL, d.AppID, route)

	req, err := http.NewRequest("POST", url, bytes.NewReader([]byte(cmd)))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bot %s", d.Token))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		log.Fatal().Msgf("%+v, %+v", resp.StatusCode, string(respBody))
	}
	return nil
}

var cmd = `
{
    "name":"coin",
    "type":1,
    "description":"Send information about a coin",
    "options":[
        {
            "name":"symbol",
			"description": "symbol to query",
            "type":3,
            "required":true,
			"min_length":3,
			"max_length": 6
        }
    ]
}
`
