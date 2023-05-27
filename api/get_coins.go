package api

import (
	"bot/entity"
	"context"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"strconv"
	"time"
)

type ListingsLatestResponse struct {
	Data []struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
		Quote  struct {
			USD struct {
				Price     float64 `json:"price"`
				Volume24H float64 `json:"volume_24h"`
				MarketCap float64 `json:"market_cap"`
			} `json:"USD"`
		} `json:"quote"`
	} `json:"data"`
	Status struct {
		Timestamp    time.Time `json:"timestamp"`
		ErrorCode    int       `json:"error_code"`
		ErrorMessage string    `json:"error_message"`
		Elapsed      int       `json:"elapsed"`
		CreditCount  int       `json:"credit_count"`
	} `json:"status"`
}

// GetCoins returns ALL coins from coinmarketcap's API
func (c clientImpl) GetCoins(ctx context.Context) ([]entity.Coin, error) {
	var coins []entity.Coin
	offset := 1
	limit := 5000
	var err error
	for {
		log.Info().Msgf("requesting coins from offset %v to %v", offset, offset+5000)
		newCoins, remaining, err := c.getCoins(context.Background(), limit, offset)
		if err != nil {
			log.Fatal().Err(err)
		}

		if remaining <= 0 {
			log.Printf("no coins were returned! sad")
			break
		}
		for _, coin := range newCoins {

			coins = append(coins, coin)
		}
		offset += 5000
	}

	return coins, err
}

func (c clientImpl) getCoins(ctx context.Context, limit, offset int) ([]entity.Coin, int, error) {
	route := "v1/cryptocurrency/listings/latest"

	url := fmt.Sprintf("%v/%v?limit=%d&start=%d", c.url, route, limit, offset)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Add(c.keyHeader, c.key)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	var jsonResp ListingsLatestResponse
	if err = json.Unmarshal(body, &jsonResp); err != nil {
		return nil, 0, err
	}

	var coins []entity.Coin

	for _, data := range jsonResp.Data {
		coin := entity.Coin{
			Id:        strconv.FormatInt(int64(data.Id), 10),
			Symbol:    data.Symbol,
			Name:      data.Name,
			MarketCap: data.Quote.USD.MarketCap,
			Price:     data.Quote.USD.Price,
			DayVolume: data.Quote.USD.Volume24H,
		}

		coins = append(coins, coin)
	}

	return coins, len(coins), nil
}
