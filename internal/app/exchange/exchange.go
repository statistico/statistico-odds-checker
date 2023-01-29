package exchange

import (
	"context"
)

type MarketFactory interface {
	CreateMarket(ctx context.Context, q *Event) (*Market, error)
}
