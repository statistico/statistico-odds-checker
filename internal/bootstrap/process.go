package bootstrap

import "github.com/statistico/statistico-odds-checker/internal/process"

func (c Container) Processor() *process.Processor {
	return process.NewProcessor(
		c.FootballEventMarketRequester(),
		c.Publisher(),
		c.Logger,
	)
}
