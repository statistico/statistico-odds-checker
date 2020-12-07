package bootstrap

import "github.com/statistico/statistico-odds-checker/internal/app/sport"

func (c Container) FootballEventMarketRequester() sport.EventMarketRequester {
	config := c.Config.FootballConfig

	return sport.NewFootballEventMarketRequester(
		c.GrpcFixtureClient(),
		c.MarketBuilder(),
		c.Logger,
		c.Clock,
		config.SupportedSeasons,
		config.Markets,
	)
}
