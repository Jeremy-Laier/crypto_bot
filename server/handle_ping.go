package server

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

func (s serverImpl) HandlePing(c echo.Context) error {
	discordPubkey, _ := hex.DecodeString(s.appPublicKey)
	verified := verify(c.Request(), ed25519.PublicKey(discordPubkey))

	if !verified {
		log.Error().Err(errors.New("failed to verify ping for some reason"))
		return c.String(http.StatusUnauthorized, "invalid request signature")
	}

	defer c.Request().Body.Close()
	var data Data
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		log.Error().Err(err)
		return c.String(http.StatusUnauthorized, "signature mismatch")
	}

	if data.Type == Ping {
		log.Info().Msg("ping found, returning 200")
		return c.JSON(200, &PingResponse{Type: 1})
	}

	log.Info().Msgf("%+v", (data.Data.Options[0].Value).(string))

	coins, err := s.db.GetCoins(context.Background(), (data.Data.Options[0].Value).(string))
	if err != nil {
		log.Error().Err(err).Msg("FAILED")
		return c.NoContent(http.StatusInternalServerError)
	}

	messageResponse := &InteractionResponse{
		Type: ChannelMessageWithSource,
		Data: &InteractionApplicationCommandCallbackData{
			Content: fmt.Sprintf("%+v", coins),
		},
	}

	var responsePayload bytes.Buffer
	err = json.NewEncoder(&responsePayload).Encode(messageResponse)
	if err != nil {
		log.Error().Err(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	resp, err := http.Post(data.ResponseURL(), "application/json", &responsePayload)
	if err != nil {
		log.Error().Err(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Error().Err(err)
			return c.NoContent(http.StatusInternalServerError)
		}

		log.Error().Err(err).Msgf("%+v", string(body))
		return c.NoContent(http.StatusInternalServerError)
	}

	log.Info().Msg("returning 204")
	return c.NoContent(http.StatusNoContent)
}

type PingResponse struct {
	Type int `json:"type"`
}
type InteractionType int

const (
	_ InteractionType = iota
	Ping
	ApplicationCommand
)

// @see https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-type
const (
	_ InteractionResponseType = iota
	Pong
	_
	_
	ChannelMessageWithSource
)

type Data struct {
	Type  InteractionType `json:"type"`
	Token string          `json:"token"`
	ID    string          `json:"id"`
	Data  struct {
		Options []ApplicationCommandInteractionDataOption `json:"options"`
		Name    string                                    `json:"name"`
		ID      string                                    `json:"id"`
	} `json:"data"`
}

func (data *Data) ResponseURL() string {
	return fmt.Sprintf("https://discord.com/api/v8/interactions/%s/%s/callback", data.ID, data.Token)
}

type ApplicationCommandInteractionDataOption struct {
	Name    string                                    `json:"name"`
	Value   interface{}                               `json:"value,omitempty"`
	Options []ApplicationCommandInteractionDataOption `json:"options,omitempty"`
}

type InteractionResponseType int

type InteractionApplicationCommandCallbackData struct {
	Content string `json:"content"`
}

type InteractionResponse struct {
	Type InteractionResponseType                    `json:"type"`
	Data *InteractionApplicationCommandCallbackData `json:"data,omitempty"`
}
