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
	"sync"
	"time"
)

type eventMarketStreamer struct {
	fixtureClient statisticodata.FixtureClient
	logger        *logrus.Logger
	clock         clockwork.Clock
	seasons       []uint64
	markets       []string
}

func (e *eventMarketStreamer) Stream(ctx context.Context, from, to time.Time, fc exchange.MarketFactory) <-chan *EventMarket {
	req := statistico.FixtureSearchRequest{
		SeasonIds:  e.seasons,
		DateBefore: &wrappers.StringValue{Value: to.Format(time.RFC3339)},
		DateAfter:  &wrappers.StringValue{Value: from.Format(time.RFC3339)},
	}

	fixtures, err := e.fixtureClient.Search(ctx, &req)

	if err != nil {
		e.logger.Errorf("Error %q fetching fixtures in football market requester", err.Error())
		return nil
	}

	ch := make(chan *EventMarket, len(fixtures))

	go e.buildEventMarkets(ctx, fixtures, ch, fc)

	return ch
}

func (e *eventMarketStreamer) buildEventMarkets(ctx context.Context, fixtures []*statistico.Fixture, ch chan<- *EventMarket, fc exchange.MarketFactory) {
	defer close(ch)
	var wg sync.WaitGroup

	for _, fx := range fixtures {
		fmt.Printf("Fetching markets for fixture %d\n", fx.Id)

		wg.Add(1)
		go e.handleFixture(ctx, fx, &wg, ch, fc)
	}

	wg.Wait()
}

func (e *eventMarketStreamer) handleFixture(ctx context.Context, f *statistico.Fixture, wg *sync.WaitGroup, ch chan<- *EventMarket, fc exchange.MarketFactory) {
	date := time.Unix(f.DateTime.Utc, 0)

	diff := date.Sub(e.clock.Now()).Minutes()

	if diff < 0 {
		wg.Done()
		return
	}

	for _, market := range e.markets {
		wg.Add(1)

		ev := exchange.Event{
			Date:   date,
			Name:   fmt.Sprintf("%s v %s", f.HomeTeam.Name, f.AwayTeam.Name),
			ID:     uint64(f.Id),
			Market: market,
		}

		go func(wg *sync.WaitGroup) {
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

				wg.Done()
				return
			}

			if m == nil || len(m.Runners) == 0 {
				wg.Done()
				return
			}

			ch <- convertToEventMarket(m, f, e.clock.Now())
			wg.Done()
		}(wg)
	}

	wg.Done()
}

func convertToEventMarket(m *exchange.Market, fix *statistico.Fixture, timestamp time.Time) *EventMarket {
	return &EventMarket{
		ID:            m.ID,
		EventID:       m.EventID,
		CompetitionID: fix.Competition.Id,
		SeasonID:      fix.Season.Id,
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
	s []uint64,
	m []string,
) EventMarketStreamer {
	return &eventMarketStreamer{
		fixtureClient: f,
		logger:        l,
		clock:         c,
		seasons:       s,
		markets:       m,
	}
}
