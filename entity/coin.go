package entity

import "fmt"

type Coin struct {
	Id        string  `json:"uuid"` // yes decode as UUID...
	Symbol    string  `json:"symbol"`
	Name      string  `json:"name"`
	MarketCap float64 `json:"marketCap"`
	Price     float64 `json:"price"`
	DayVolume float64 `json:"24hVolume"`
}

func (c Coin) String() string {
	return fmt.Sprintf("\nName: %v\nMarket Cap: %f\nPrice($USD): %f\nVolume 24 Hours: %f\n", c.Name, c.MarketCap, c.Price, c.DayVolume)
}
