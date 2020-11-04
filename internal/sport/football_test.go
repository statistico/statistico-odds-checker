package sport_test

import (
	"context"
	"errors"
	"github.com/jonboulle/clockwork"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-odds-checker/internal/exchange"
	"github.com/statistico/statistico-odds-checker/internal/grpc"
	"github.com/statistico/statistico-odds-checker/internal/grpc/proto"
	"github.com/statistico/statistico-odds-checker/internal/market"
	"github.com/statistico/statistico-odds-checker/internal/sport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestFootballEventMarketRequester_FindEventMarkets(t *testing.T) {
	t.Run("parses fixtures and markets and pushes them into the provided channel", func(t *testing.T) {
		t.Helper()

		fixClient := new(grpc.MockFixtureGrpcClient)
		builder := new(market.MockMarketBuilder)
		logger, _ := test.NewNullLogger()
		clock := clockwork.NewFakeClockAt(time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC))
		seasons := []uint64{1234, 5678}
		markets := []string{"OVER_UNDER_25", "1X2"}

		r := sport.NewFootballEventMarketRequester(fixClient, builder, logger, clock, seasons, markets)

		ctx := context.Background()

		fixReq := mock.MatchedBy(func(r *proto.FixtureSearchRequest) bool {
			assert.Equal(t, []uint64{1234, 5678}, r.SeasonIds)
			assert.Equal(t, "2019-01-14T10:00:00Z", r.DateAfter.GetValue())
			assert.Equal(t, "2019-01-14T12:00:00Z", r.DateBefore.GetValue())
			return true
		})

		fixtures := []*proto.Fixture{
			{
				Id: 349811,
				HomeTeam: &proto.Team{
					Name: "West Ham United",
				},
				AwayTeam: &proto.Team{
					Name: "Arsenal",
				},
				DateTime: &proto.Date{
					Utc: 1547465400,
				},
			},
		}

		fixClient.On("Search", ctx, fixReq).Return(fixtures, nil)

		bq := mock.MatchedBy(func(q *market.BuilderQuery) bool {
			assert.Equal(t, time.Unix(1547465400, 0), q.Date)
			assert.Equal(t, "West Ham United v Arsenal", q.Event)
			assert.Equal(t, uint64(349811), q.EventID)
			assert.Equal(t, "football", q.Sport)
			assert.Equal(t, []string{"OVER_UNDER_25", "1X2"}, q.Markets)
			return true
		})

		marketSlice := []*market.Market{
			{
				ID:             "1.254912",
				EventID:        349811,
				Name:           "OVER_UNDER_25",
				Side:           "BACK",
				Exchange:       "betfair",
				ExchangeRunners: []*exchange.Runner{},
				StatisticoOdds: []*proto.Odds{},
			},
			{
				ID:             "1.3410292",
				EventID:        349811,
				Name:           "1X2",
				Side:           "BACK",
				Exchange:       "betfair",
				ExchangeRunners: []*exchange.Runner{},
				StatisticoOdds: []*proto.Odds{},
			},
		}

		builder.On("Build", ctx, bq).Return(marketChannel(marketSlice))

		from := time.Date(2019, 01, 14, 10, 00, 00, 00, time.UTC)
		to := time.Date(2019, 01, 14, 12, 00, 00, 00, time.UTC)

		ch := r.FindEventMarkets(ctx, from, to)

		emOne := <-ch
		emTwo := <-ch

		a := assert.New(t)

		a.Equal("1.254912", emOne.ID)
		a.Equal(uint64(349811), emOne.EventID)
		a.Equal("football", emOne.Sport)
		a.Equal(int64(1547465400), emOne.EventDate)
		a.Equal("OVER_UNDER_25", emOne.MarketName)
		a.Equal("BACK", emOne.Side)
		a.Equal("betfair", emOne.Exchange)
		a.Equal([]*exchange.Runner{}, emOne.ExchangeRunners)
		a.Equal([]*proto.Odds{}, emOne.StatisticoOdds)
		a.Equal(int64(1547465100), emOne.Timestamp)

		a.Equal("1.3410292", emTwo.ID)
		a.Equal(uint64(349811), emTwo.EventID)
		a.Equal("football", emTwo.Sport)
		a.Equal(int64(1547465400), emTwo.EventDate)
		a.Equal("1X2", emTwo.MarketName)
		a.Equal("BACK", emTwo.Side)
		a.Equal("betfair", emTwo.Exchange)
		a.Equal([]*exchange.Runner{}, emTwo.ExchangeRunners)
		a.Equal([]*proto.Odds{}, emTwo.StatisticoOdds)
		a.Equal(int64(1547465100), emTwo.Timestamp)
	})

	t.Run("returns nil if error fetching fixtures via fixture client", func(t *testing.T) {
		t.Helper()

		fixClient := new(grpc.MockFixtureGrpcClient)
		builder := new(market.MockMarketBuilder)
		logger, hook := test.NewNullLogger()
		clock := clockwork.NewFakeClockAt(time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC))
		seasons := []uint64{1234, 5678}
		markets := []string{"OVER_UNDER_25", "1X2"}

		r := sport.NewFootballEventMarketRequester(fixClient, builder, logger, clock, seasons, markets)

		fixReq := mock.MatchedBy(func(r *proto.FixtureSearchRequest) bool {
			assert.Equal(t, []uint64{1234, 5678}, r.SeasonIds)
			assert.Equal(t, "2019-01-14T10:00:00Z", r.DateAfter.GetValue())
			assert.Equal(t, "2019-01-14T12:00:00Z", r.DateBefore.GetValue())
			return true
		})

		ctx := context.Background()

		fixClient.On("Search", ctx, fixReq).Return([]*proto.Fixture{}, errors.New("error fetching fixtures"))

		from := time.Date(2019, 01, 14, 10, 00, 00, 00, time.UTC)
		to := time.Date(2019, 01, 14, 12, 00, 00, 00, time.UTC)

		ch := r.FindEventMarkets(ctx, from, to)

		assert.Nil(t, ch)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Error \"error fetching fixtures\" fetching fixtures in football market requester", hook.LastEntry().Message)

		builder.AssertNotCalled(t, "Build")
	})
}

func marketChannel(markets []*market.Market) <-chan *market.Market {
	ch := make(chan *market.Market, len(markets))

	for _, m := range markets {
		ch <- m
	}

	close(ch)

	return ch
}
