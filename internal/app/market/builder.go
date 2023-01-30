package market

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"sync"
	"time"
)

type Builder interface {
	Build(ctx context.Context, q *BuilderQuery) <-chan *Market
}

type BuilderQuery struct {
	Date    time.Time
	Event   string
	EventID uint64
	Sport   string
	Markets []string
}

type builder struct {
	exchange exchange.MarketFactory
	logger   *logrus.Logger
}

func (b *builder) Build(ctx context.Context, q *BuilderQuery) <-chan *Market {
	ch := make(chan *Market, 100)

	go b.parseQuery(ctx, q, ch)

	return ch
}

func (b *builder) parseQuery(ctx context.Context, q *BuilderQuery, ch chan<- *Market) {
	defer close(ch)
	var wg sync.WaitGroup

	for _, market := range q.Markets {
		wg.Add(1)
		go b.buildMarket(ctx, q, market, &wg, ch)
	}

	wg.Wait()
}

func (b *builder) buildMarket(ctx context.Context, q *BuilderQuery, market string, wg *sync.WaitGroup, ch chan<- *Market) {
	bm, err := b.fetchExchangeMarket(ctx, q, market)

	if err != nil {
		b.logError(err, q.EventID, market)
		wg.Done()
		return
	}

	m := Market{
		ID:              bm.ID,
		EventID:         q.EventID,
		Name:            bm.Name,
		Exchange:        bm.ExchangeName,
		ExchangeRunners: bm.Runners,
	}

	ch <- &m

	wg.Done()
}

func (b *builder) fetchExchangeMarket(ctx context.Context, q *BuilderQuery, market string) (*exchange.Market, error) {
	eq := exchange.Event{
		Name:   q.Event,
		Date:   q.Date,
		Market: market,
	}

	mk, err := b.exchange.CreateMarket(ctx, &eq)

	if err != nil {
		return nil, err
	}

	return mk, nil
}

func (b *builder) logError(e error, eventID uint64, market string) {
	switch e.(type) {
	case *exchange.NoEventMarketError:
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

func NewBuilder(m exchange.MarketFactory, l *logrus.Logger) Builder {
	return &builder{
		exchange: m,
		logger:   l,
	}
}
