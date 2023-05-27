package api

import (
	"bot/entity"
	"context"
)

type Client interface {
	GetCoins(ctx context.Context) ([]entity.Coin, error)
}
