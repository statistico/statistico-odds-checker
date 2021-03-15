package sport

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/jonboulle/clockwork"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-data-go-grpc-client"
	"github.com/statistico/statistico-odds-checker/internal/app/market"
	"github.com/statistico/statistico-proto/go"
	"time"
)

const football = "football"

type footballEventMarketRequester struct {
	fixtureClient statisticodata.FixtureClient
	builder       market.Builder
	logger        *logrus.Logger
	clock         clockwork.Clock
	seasons       []uint64
	markets       []string
}

func (f *footballEventMarketRequester) FindEventMarkets(ctx context.Context, from, to time.Time) <-chan *EventMarket {
	req := statistico.FixtureSearchRequest{
		SeasonIds:  f.seasons,
		DateBefore: &wrappers.StringValue{Value: to.Format(time.RFC3339)},
		DateAfter:  &wrappers.StringValue{Value: from.Format(time.RFC3339)},
	}

	fixtures, err := f.fixtureClient.Search(ctx, &req)

	if err != nil {
		f.logger.Errorf("Error %q fetching fixtures in football market requester", err.Error())
		return nil
	}

	ch := make(chan *EventMarket, len(fixtures))

	go f.buildEventMarkets(ctx, fixtures, ch)

	return ch
}

func (f *footballEventMarketRequester) buildEventMarkets(ctx context.Context, fixtures []*statistico.Fixture, ch chan<- *EventMarket) {
	for _, fx := range fixtures {
		date := time.Unix(fx.DateTime.Utc, 0)

		diff := date.Sub(f.clock.Now()).Hours()

		if diff >= 3 || diff < 0 {
			continue
		}

		q := market.BuilderQuery{
			Date:    date,
			Event:   fmt.Sprintf("%s v %s", fx.HomeTeam.Name, fx.AwayTeam.Name),
			EventID: uint64(fx.Id),
			Sport:   football,
			Markets: f.markets,
		}

		for m := range f.builder.Build(ctx, &q) {
			ch <- convertToEventMarket(m, fx, f.clock.Now())
		}
	}

	close(ch)
}

func convertToEventMarket(m *market.Market, fix *statistico.Fixture, timestamp time.Time) *EventMarket {
	return &EventMarket{
		ID:            m.ID,
		EventID:       m.EventID,
		CompetitionID: fix.Competition.Id,
		SeasonID:      fix.Season.Id,
		Sport:         football,
		EventDate:     fix.DateTime.Rfc,
		MarketName:    m.Name,
		Exchange:      m.Exchange,
		Runners:       m.ExchangeRunners,
		Timestamp:     timestamp.Unix(),
	}
}

func NewFootballEventMarketRequester(
	f statisticodata.FixtureClient,
	b market.Builder,
	l *logrus.Logger,
	c clockwork.Clock,
	s []uint64,
	m []string,
) EventMarketRequester {
	return &footballEventMarketRequester{
		fixtureClient: f,
		builder:       b,
		logger:        l,
		clock:         c,
		seasons:       s,
		markets:       m,
	}
}
