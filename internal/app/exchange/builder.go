package exchange

import (
	"context"
	"github.com/sirupsen/logrus"
	"sync"
)

type MarketBuilder interface {
	Build(ctx context.Context, e *Event) <-chan *Market
}

type marketBuilder struct {
	factories []MarketFactory
	logger    *logrus.Logger
}

func (b *marketBuilder) Build(ctx context.Context, e *Event) <-chan *Market {
	ch := make(chan *Market, 100)

	go b.handleEvent(ctx, e, ch)

	return ch
}

func (b *marketBuilder) handleEvent(ctx context.Context, e *Event, ch chan<- *Market) {
	defer close(ch)
	var wg sync.WaitGroup

	for _, f := range b.factories {
		wg.Add(1)
		go b.buildMarkets(ctx, e, f, &wg, ch)
	}

	wg.Wait()
}

func (b *marketBuilder) buildMarkets(ctx context.Context, e *Event, factory MarketFactory, wg *sync.WaitGroup, ch chan<- *Market) {
	mk, err := factory.CreateMarket(ctx, e)

	if err != nil {
		b.logError(err, e.ID, e.Market)
		wg.Done()
		return
	}

	if len(mk.Runners) == 0 {
		wg.Done()
		return
	}

	ch <- mk

	wg.Done()
}

func (b *marketBuilder) logError(e error, eventID uint64, market string) {
	switch e.(type) {
	case *NoEventMarketError:
		return
	default:
		b.logger.Errorf(
			"Error when calling client '%s' for event %d and market %s",
			e.Error(),
			eventID,
			market,
		)
	}
}

func NewMarketBuilder(f []MarketFactory, l *logrus.Logger) MarketBuilder {
	return &marketBuilder{
		factories: f,
		logger:    l,
	}
}
