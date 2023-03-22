package process

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-odds-checker/internal/app/publish"
	"github.com/statistico/statistico-odds-checker/internal/app/stream"
	"time"
)

type Processor struct {
	streamer  stream.EventMarketStreamer
	resolver  exchange.MarketFactoryResolver
	publisher publish.Publisher
	logger    *logrus.Logger
}

func (p *Processor) Process(ctx context.Context, from, to time.Time, exchange string) error {
	f, err := p.resolver.Resolve(exchange)

	if err != nil {
		return err
	}

	markets := p.streamer.Stream(ctx, from, to, f)

	for m := range markets {
		err := p.publisher.PublishMarket(m)

		if err != nil {
			p.logger.Errorf("Error publishing market %q", err)
		}
	}

	return nil
}

func NewProcessor(s stream.EventMarketStreamer, r exchange.MarketFactoryResolver, p publish.Publisher, l *logrus.Logger) *Processor {
	return &Processor{
		streamer:  s,
		resolver:  r,
		publisher: p,
		logger:    l,
	}
}
