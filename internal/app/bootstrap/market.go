package bootstrap

import "github.com/statistico/statistico-odds-checker/internal/app/market"

func (c Container) MarketBuilder() market.Builder {
	return market.NewBuilder(
		c.MarketFactory(),
		c.Logger,
	)
}
