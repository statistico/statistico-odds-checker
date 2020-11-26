package bootstrap

import "github.com/statistico/statistico-odds-checker/internal/market"

func (c Container) MarketBuilder() market.Builder {
	return market.NewBuilder(
		c.MarketRequester(),
		c.Logger,
	)
}
