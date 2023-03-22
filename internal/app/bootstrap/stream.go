package bootstrap

import (
	"github.com/statistico/statistico-odds-checker/internal/app/stream"
)

func (c Container) EventMarketStreamer() stream.EventMarketStreamer {
	config := c.Config.FootballConfig

	return stream.NewEventMarketStreamer(
		c.DataServiceResultClient(),
		c.Logger,
		c.Clock,
		config.SupportedSeasons,
		config.Markets,
	)
}
