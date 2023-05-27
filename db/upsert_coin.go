package db

import (
	"bot/entity"
	"context"
	"errors"
	"fmt"
)

func (d dbImpl) UpsertCoin(ctx context.Context, coin entity.Coin) error {
	query := `
		INSERT INTO coin
		(id, symbol, name, market_cap, price, day_volume)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id, symbol)
		DO UPDATE SET
			market_cap = excluded.market_cap,
			price = excluded.price,
			day_volume = excluded.day_volume
		;
	`
	result, err := d.ExecContext(ctx, query, coin.Id, coin.Symbol, coin.Name, coin.MarketCap, coin.Price, coin.DayVolume)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New(fmt.Sprintf("failed to insert coin ID %v symbol %v", coin.Id, coin.Symbol))
	}

	return nil
}
