package process

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-checker/internal/app/publish"
	"github.com/statistico/statistico-odds-checker/internal/app/stream"
	"time"
)

type Processor struct {
	streamer  stream.EventMarketStreamer
	publisher publish.Publisher
	logger    *logrus.Logger
}

func (p *Processor) Process(ctx context.Context, from, to time.Time) error {
	var markets <-chan *stream.EventMarket

	markets = p.streamer.Stream(ctx, from, to)

	for m := range markets {
		err := p.publisher.PublishMarket(m)

		if err != nil {
			p.logger.Errorf("Error publishing market %q", err)
		}
	}

	return nil
}

func NewProcessor(s stream.EventMarketStreamer, p publish.Publisher, l *logrus.Logger) *Processor {
	return &Processor{
		streamer:  s,
		publisher: p,
		logger:    l,
	}
}
