package stream_test

import (
	"context"
	"errors"
	"github.com/jonboulle/clockwork"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-odds-checker/internal/app/stream"
	"github.com/statistico/statistico-proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestEventMarketStreamer_Stream(t *testing.T) {
	t.Run("parses fixtures and markets and pushes them into the provided channel", func(t *testing.T) {
		t.Helper()

		fixClient := new(MockFixtureClient)
		builder := new(exchange.MockMarketBuilder)
		logger, _ := test.NewNullLogger()
		clock := clockwork.NewFakeClockAt(time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC))
		seasons := []uint64{1234, 5678}
		markets := []string{"OVER_UNDER_25", "1X2"}

		st := stream.NewEventMarketStreamer(fixClient, builder, logger, clock, seasons, markets)

		ctx := context.Background()

		fixReq := mock.MatchedBy(func(r *statistico.FixtureSearchRequest) bool {
			assert.Equal(t, []uint64{1234, 5678}, r.SeasonIds)
			assert.Equal(t, "2019-01-14T10:00:00Z", r.DateAfter.GetValue())
			assert.Equal(t, "2019-01-14T12:00:00Z", r.DateBefore.GetValue())
			return true
		})

		fixtures := []*statistico.Fixture{
			{
				Id: 349811,
				HomeTeam: &statistico.Team{
					Name: "West Ham United",
				},
				AwayTeam: &statistico.Team{
					Name: "Arsenal",
				},
				DateTime: &statistico.Date{
					Utc: 1547465400,
					Rfc: "2019-01-14T12:00:00Z",
				},
				Competition: &statistico.Competition{Id: 8},
				Season:      &statistico.Season{Id: 17420},
			},
		}

		fixClient.On("Search", ctx, fixReq).Return(fixtures, nil)

		event := mock.MatchedBy(func(e *exchange.Event) bool {
			if e.Market == "1X2" {
				assert.NotEqual(t, "OVER_UNDER_25", e.Market)
			}

			if e.Market == "OVER_UNDER_25" {
				assert.NotEqual(t, "1X2", e.Market)
			}

			assert.Equal(t, time.Unix(1547465400, 0), e.Date)
			assert.Equal(t, "West Ham United v Arsenal", e.Name)
			assert.Equal(t, uint64(349811), e.ID)
			return true
		})

		marketSlice := []*exchange.Market{
			{
				ID:       "1.254912",
				EventID:  349811,
				Name:     "OVER_UNDER_25",
				Exchange: "BETFAIR",
				Runners:  []*exchange.Runner{},
			},
			{
				ID:       "1.3410292",
				EventID:  349811,
				Name:     "MATCH_ODDS",
				Exchange: "PINNACLE",
				Runners:  []*exchange.Runner{},
			},
		}

		builder.On("Build", ctx, event).Return(marketChannel(marketSlice)).Once()
		builder.On("Build", ctx, event).Return(marketChannel([]*exchange.Market{})).Once()

		from := time.Date(2019, 01, 14, 10, 00, 00, 00, time.UTC)
		to := time.Date(2019, 01, 14, 12, 00, 00, 00, time.UTC)

		ch := st.Stream(ctx, from, to)

		emOne := <-ch
		emTwo := <-ch

		<-ch

		a := assert.New(t)

		a.Equal("1.254912", emOne.ID)
		a.Equal(uint64(349811), emOne.EventID)
		a.Equal(uint64(8), emOne.CompetitionID)
		a.Equal(uint64(17420), emOne.SeasonID)
		a.Equal(int64(1547465400), emOne.EventDate)
		a.Equal("OVER_UNDER_25", emOne.MarketName)
		a.Equal("BETFAIR", emOne.Exchange)
		a.Equal([]*exchange.Runner{}, emOne.Runners)
		a.Equal(int64(1547465100), emOne.Timestamp)

		a.Equal("1.3410292", emTwo.ID)
		a.Equal(uint64(349811), emTwo.EventID)
		a.Equal(uint64(8), emOne.CompetitionID)
		a.Equal(uint64(17420), emOne.SeasonID)
		a.Equal(int64(1547465400), emTwo.EventDate)
		a.Equal("MATCH_ODDS", emTwo.MarketName)
		a.Equal("PINNACLE", emTwo.Exchange)
		a.Equal([]*exchange.Runner{}, emTwo.Runners)
		a.Equal(int64(1547465100), emTwo.Timestamp)

		builder.AssertExpectations(t)
		builder.AssertNumberOfCalls(t, "Build", 2)
	})

	t.Run("builder is not called if event date and current date difference is greater than two hours", func(t *testing.T) {
		t.Helper()

		fixClient := new(MockFixtureClient)
		builder := new(exchange.MockMarketBuilder)
		logger, _ := test.NewNullLogger()
		clock := clockwork.NewFakeClockAt(time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC))
		seasons := []uint64{1234, 5678}
		markets := []string{"OVER_UNDER_25", "1X2"}

		st := stream.NewEventMarketStreamer(fixClient, builder, logger, clock, seasons, markets)

		fixReq := mock.MatchedBy(func(r *statistico.FixtureSearchRequest) bool {
			assert.Equal(t, []uint64{1234, 5678}, r.SeasonIds)
			assert.Equal(t, "2019-01-14T10:00:00Z", r.DateAfter.GetValue())
			assert.Equal(t, "2019-01-14T12:00:00Z", r.DateBefore.GetValue())
			return true
		})

		fixtures := []*statistico.Fixture{
			{
				Id: 349811,
				HomeTeam: &statistico.Team{
					Name: "West Ham United",
				},
				AwayTeam: &statistico.Team{
					Name: "Arsenal",
				},
				DateTime: &statistico.Date{
					Utc: 1547496000,
					Rfc: "2019-01-14T20:00:00Z",
				},
				Competition: &statistico.Competition{Id: 8},
				Season:      &statistico.Season{Id: 17420},
			},
		}

		ctx := context.Background()
		from := time.Date(2019, 01, 14, 10, 00, 00, 00, time.UTC)
		to := time.Date(2019, 01, 14, 12, 00, 00, 00, time.UTC)

		fixClient.On("Search", ctx, fixReq).Return(fixtures, nil)

		builder.AssertNotCalled(t, "Build")

		ch := st.Stream(ctx, from, to)

		one := <-ch

		assert.Nil(t, one)
	})

	t.Run("build is not called if event date and current date difference is less than zero", func(t *testing.T) {
		t.Helper()

		fixClient := new(MockFixtureClient)
		builder := new(exchange.MockMarketBuilder)
		logger, _ := test.NewNullLogger()
		clock := clockwork.NewFakeClockAt(time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC))
		seasons := []uint64{1234, 5678}
		markets := []string{"OVER_UNDER_25", "1X2"}

		st := stream.NewEventMarketStreamer(fixClient, builder, logger, clock, seasons, markets)

		fixReq := mock.MatchedBy(func(r *statistico.FixtureSearchRequest) bool {
			assert.Equal(t, []uint64{1234, 5678}, r.SeasonIds)
			assert.Equal(t, "2019-01-14T10:00:00Z", r.DateAfter.GetValue())
			assert.Equal(t, "2019-01-14T12:00:00Z", r.DateBefore.GetValue())
			return true
		})

		fixtures := []*statistico.Fixture{
			{
				Id: 349811,
				HomeTeam: &statistico.Team{
					Name: "West Ham United",
				},
				AwayTeam: &statistico.Team{
					Name: "Arsenal",
				},
				DateTime: &statistico.Date{
					Utc: 1547496000,
					Rfc: "2019-01-14T10:00:00Z",
				},
				Competition: &statistico.Competition{Id: 8},
				Season:      &statistico.Season{Id: 17420},
			},
		}

		ctx := context.Background()
		from := time.Date(2019, 01, 14, 10, 00, 00, 00, time.UTC)
		to := time.Date(2019, 01, 14, 12, 00, 00, 00, time.UTC)

		fixClient.On("Search", ctx, fixReq).Return(fixtures, nil)

		builder.AssertNotCalled(t, "Build")

		ch := st.Stream(ctx, from, to)

		one := <-ch

		assert.Nil(t, one)
	})

	t.Run("returns nil if error fetching fixtures via event client", func(t *testing.T) {
		t.Helper()

		fixClient := new(MockFixtureClient)
		builder := new(exchange.MockMarketBuilder)
		logger, hook := test.NewNullLogger()
		clock := clockwork.NewFakeClockAt(time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC))
		seasons := []uint64{1234, 5678}
		markets := []string{"OVER_UNDER_25", "1X2"}

		st := stream.NewEventMarketStreamer(fixClient, builder, logger, clock, seasons, markets)

		fixReq := mock.MatchedBy(func(r *statistico.FixtureSearchRequest) bool {
			assert.Equal(t, []uint64{1234, 5678}, r.SeasonIds)
			assert.Equal(t, "2019-01-14T10:00:00Z", r.DateAfter.GetValue())
			assert.Equal(t, "2019-01-14T12:00:00Z", r.DateBefore.GetValue())
			return true
		})

		ctx := context.Background()

		fixClient.On("Search", ctx, fixReq).Return([]*statistico.Fixture{}, errors.New("error fetching fixtures"))

		from := time.Date(2019, 01, 14, 10, 00, 00, 00, time.UTC)
		to := time.Date(2019, 01, 14, 12, 00, 00, 00, time.UTC)

		ch := st.Stream(ctx, from, to)

		assert.Nil(t, ch)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Error \"error fetching fixtures\" fetching fixtures in football market requester", hook.LastEntry().Message)

		builder.AssertNotCalled(t, "Build")
	})
}

func marketChannel(markets []*exchange.Market) <-chan *exchange.Market {
	ch := make(chan *exchange.Market, len(markets))

	for _, m := range markets {
		ch <- m
	}

	close(ch)

	return ch
}

type MockFixtureClient struct {
	mock.Mock
}

func (m *MockFixtureClient) Search(ctx context.Context, req *statistico.FixtureSearchRequest) ([]*statistico.Fixture, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]*statistico.Fixture), args.Error(1)
}

func (m *MockFixtureClient) ByID(ctx context.Context, fixtureID uint64) (*statistico.Fixture, error) {
	args := m.Called(ctx, fixtureID)
	return args.Get(0).(*statistico.Fixture), args.Error(1)
}
