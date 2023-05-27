package db

import (
	"bot/entity"
	"context"
)

type DB interface {
	GetCoins(ctx context.Context, symbol string) ([]entity.Coin, error)
	UpsertCoin(ctx context.Context, coin entity.Coin) error
}
