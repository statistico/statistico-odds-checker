package exchange

import "context"

type MarketFactory interface {
	CreateMarket(ctx context.Context, e *Event) (*Market, error)
}
