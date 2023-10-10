package stream

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/jonboulle/clockwork"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-football-data-go-grpc-client"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-proto/go"
	"strconv"
	"sync"
	"time"
)

type eventMarketStreamer struct {
	fixtureClient statisticodata.FixtureClient
	logger        *logrus.Logger
	clock         clockwork.Clock
	markets       []string
}

func (e *eventMarketStreamer) Stream(ctx context.Context, from, to time.Time, fc exchange.MarketFactory, market string) <-chan *EventMarket {
	req := statistico.FixtureSearchRequest{
		DateBefore: &wrappers.StringValue{Value: to.Format(time.RFC3339)},
		DateAfter:  &wrappers.StringValue{Value: from.Format(time.RFC3339)},
	}

	fixtures, err := e.fixtureClient.Search(ctx, &req)

	if err != nil {
		e.logger.Errorf("Error %q fetching fixtures in football market requester", err.Error())
		return nil
	}

	ch := make(chan *EventMarket, len(fixtures))

	go e.buildEventMarkets(ctx, fixtures, ch, fc, market)

	return ch
}

func (e *eventMarketStreamer) buildEventMarkets(ctx context.Context, fixtures []*statistico.Fixture, ch chan<- *EventMarket, fc exchange.MarketFactory, market string) {
	defer close(ch)
	var wg sync.WaitGroup

	for _, fx := range fixtures {
		fmt.Printf("Fetching markets for fixture %d\n", fx.Id)

		wg.Add(1)
		go e.handleFixture(ctx, fx, &wg, ch, fc, market)
	}

	wg.Wait()
}

func (e *eventMarketStreamer) handleFixture(ctx context.Context, f *statistico.Fixture, wg *sync.WaitGroup, ch chan<- *EventMarket, fc exchange.MarketFactory, market string) {
	date := time.Unix(f.DateTime.Utc, 0)

	diff := date.Sub(e.clock.Now()).Minutes()

	if diff < 0 {
		wg.Done()
		return
	}

	markets := e.markets

	if market != "" {
		markets = []string{market}
	}

	for _, market := range markets {
		ev := exchange.Event{
			Date:   date,
			Name:   fmt.Sprintf("%s v %s", f.HomeTeam.Name, f.AwayTeam.Name),
			ID:     uint64(f.Id),
			Market: market,
		}

		m, err := fc.CreateMarket(ctx, &ev)

		if err != nil {
			switch err.(type) {
			case *exchange.NoEventMarketError:
				e.logger.Info(err.Error())
				break
			default:
				e.logger.Errorf(
					"error when calling factory '%s' for event %d and market %s and exchange %s",
					err.Error(),
					ev.ID,
					ev.Market,
					fc.Exchange(),
				)
				break
			}

			continue
		}

		if m == nil || len(m.Runners) == 0 {
			continue
		}

		ch <- convertToEventMarket(m, f, e.clock.Now())
	}

	wg.Done()
}

func convertToEventMarket(m *exchange.Market, fix *statistico.Fixture, timestamp time.Time) *EventMarket {
	var r uint64

	if fix.Round != nil {
		r, _ = strconv.ParseUint(fix.Round.Name, 10, 64)
	}

	return &EventMarket{
		ID:            m.ID,
		EventID:       m.EventID,
		CompetitionID: fix.Competition.Id,
		SeasonID:      fix.Season.Id,
		Round:         r,
		EventDate:     fix.DateTime.Utc,
		MarketName:    m.Name,
		Exchange:      m.Exchange,
		Runners:       m.Runners,
		Timestamp:     timestamp.Unix(),
	}
}

func NewEventMarketStreamer(
	f statisticodata.FixtureClient,
	l *logrus.Logger,
	c clockwork.Clock,
	m []string,
) EventMarketStreamer {
	return &eventMarketStreamer{
		fixtureClient: f,
		logger:        l,
		clock:         c,
		markets:       m,
	}
}
