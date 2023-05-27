package api

import "os"

type clientImpl struct {
	key       string
	keyHeader string
	url       string
}

func New() Client {
	return &clientImpl{
		key:       os.Getenv("COINMARKET_API_KEY"),
		keyHeader: "X-CMC_PRO_API_KEY",
		url:       "https://pro-api.coinmarketcap.com",
	}
}
