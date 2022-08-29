package process

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-checker/internal/app/publish"
	sp "github.com/statistico/statistico-odds-checker/internal/app/sport"
	"time"
)

const Football = "FOOTBALL"

type Processor struct {
	football  sp.EventMarketRequester
	publisher publish.Publisher
	logger    *logrus.Logger
}

func (p *Processor) Process(ctx context.Context, sport string, from, to time.Time) error {
	var markets <-chan *sp.EventMarket

	switch sport {
	case Football:
		markets = p.football.FindEventMarkets(ctx, from, to)
		break
	default:
		return errors.New("sport provided is not supported")
	}

	for m := range markets {
		err := p.publisher.PublishMarket(m)

		if err != nil {
			p.logger.Errorf("Error publishing market %q", err)
		}
	}

	return nil
}

func NewProcessor(f sp.EventMarketRequester, p publish.Publisher, l *logrus.Logger) *Processor {
	return &Processor{
		football:  f,
		publisher: p,
		logger:    l,
	}
}
