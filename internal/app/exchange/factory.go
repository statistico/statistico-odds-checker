package exchange

import "context"

type MarketFactory interface {
	// CreateMarket receives on context.Context struct and an Event struct. Implementations of this interface
	// will internally build a Market struct based on the provided Event parameters.
	CreateMarket(ctx context.Context, e *Event) (*Market, error)
}
