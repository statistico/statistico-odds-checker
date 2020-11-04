package market

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-odds-checker/internal/exchange"
	"github.com/statistico/statistico-odds-checker/internal/grpc"
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
	compilerClient grpc.OddsCompilerClient
	exchange       exchange.MarketRequester
	logger         *logrus.Logger
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
	odds, err := b.compilerClient.GetEventMarket(ctx, q.EventID, market)

	if err != nil {
		b.logError(err, q.EventID, market)
		wg.Done()
		return
	}

	bm, err := b.fetchExchangeMarket(ctx, q, market)

	if err != nil {
		b.logError(err, q.EventID, market)
		wg.Done()
		return
	}

	m := Market{
		ID: bm.ID,
		EventID:        q.EventID,
		Name:           market,
		Exchange:       bm.ExchangeName,
		Side:          	bm.Side,
		ExchangeRunners: bm.Runners,
		StatisticoOdds: odds,
	}

	ch <- &m

	wg.Done()
}

func (b *builder) fetchExchangeMarket(ctx context.Context, q *BuilderQuery, market string) (*exchange.Market, error) {
	eq := exchange.Query{
		Event:  q.Event,
		Date:   q.Date,
		Market: market,
		Sport:  q.Sport,
	}

	mk, err := b.exchange.Fetch(ctx, &eq)

	if err != nil {
		return nil, err
	}

	return mk, nil
}

func (b *builder) logError(e error, eventID uint64, market string) {
	b.logger.Errorf("Error when calling client '%s' for event %d and market %s", e.Error(), eventID, market)
}

func NewBuilder(o grpc.OddsCompilerClient, m exchange.MarketRequester, l *logrus.Logger) Builder {
	return &builder{
		compilerClient: o,
		exchange:       m,
		logger:         l,
	}
}
