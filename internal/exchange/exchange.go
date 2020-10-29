package exchange

import (
	"context"
)

type MarketRequester interface {
	Fetch(ctx context.Context, q *Query) (*Market, error)
}
