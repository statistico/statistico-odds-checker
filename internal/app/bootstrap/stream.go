package bootstrap

import (
	"github.com/statistico/statistico-odds-checker/internal/app/stream"
)

func (c Container) EventMarketStreamer() stream.EventMarketStreamer {
	return stream.NewEventMarketStreamer(
		c.DataServiceResultClient(),
		c.Logger,
		c.Clock,
	)
}
