package db

import (
	"bot/entity"
	"context"
	"database/sql"
)

func (d dbImpl) GetCoins(ctx context.Context, symbol string) ([]entity.Coin, error) {
	query := `
		SELECT symbol, name, market_cap, price, day_volume 
		FROM coin
	`

	var rows *sql.Rows
	var err error

	if symbol != "" {
		query = query + `WHERE symbol = $1`
		rows, err = d.QueryContext(ctx, query, symbol)
	} else {
		rows, err = d.QueryContext(ctx, query)
	}

	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	defer rows.Close()

	var coins []entity.Coin

	for rows.Next() {
		var coin entity.Coin
		err = rows.Scan(&coin.Symbol, &coin.Name, &coin.MarketCap, &coin.Price, &coin.DayVolume)
		if err != nil {
			return nil, err
		}

		coins = append(coins, coin)
	}

	return coins, nil
}
