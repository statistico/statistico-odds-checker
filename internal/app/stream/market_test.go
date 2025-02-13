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
		factory := new(exchange.MockMarketFactory)
		logger, _ := test.NewNullLogger()
		clock := clockwork.NewFakeClockAt(time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC))
		markets := []string{"OVER_UNDER_25", "1X2"}

		st := stream.NewEventMarketStreamer(fixClient, logger, clock, markets)

		ctx := context.Background()

		fixReq := mock.MatchedBy(func(r *statistico.FixtureSearchRequest) bool {
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
				Round:       &statistico.Round{Name: "5"},
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

		marketOne := exchange.Market{
			ID:       "1.254912",
			EventID:  349811,
			Name:     "OVER_UNDER_25",
			Exchange: "BETFAIR",
			Runners: []*exchange.Runner{
				{
					ID:         "0",
					Name:       "OVER",
					BackPrices: nil,
					LayPrices:  nil,
				},
			},
		}

		factory.On("CreateMarket", ctx, event).Return(&marketOne, nil)

		from := time.Date(2019, 01, 14, 10, 00, 00, 00, time.UTC)
		to := time.Date(2019, 01, 14, 12, 00, 00, 00, time.UTC)

		ch := st.Stream(ctx, from, to, factory, "")

		emOne := <-ch

		<-ch

		a := assert.New(t)

		a.Equal("1.254912", emOne.ID)
		a.Equal(uint64(349811), emOne.EventID)
		a.Equal(uint64(8), emOne.CompetitionID)
		a.Equal(uint64(17420), emOne.SeasonID)
		a.Equal(uint64(5), emOne.Round)
		a.Equal(int64(1547465400), emOne.EventDate)
		a.Equal("OVER_UNDER_25", emOne.MarketName)
		a.Equal("BETFAIR", emOne.Exchange)
		a.Equal(marketOne.Runners, emOne.Runners)
		a.Equal(int64(1547465100), emOne.Timestamp)

		fixClient.AssertExpectations(t)
		factory.AssertExpectations(t)
	})

	t.Run("parses fixtures and markets and pushes them into the provided channel by overriding default markets", func(t *testing.T) {
		t.Helper()

		fixClient := new(MockFixtureClient)
		factory := new(exchange.MockMarketFactory)
		logger, _ := test.NewNullLogger()
		clock := clockwork.NewFakeClockAt(time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC))
		markets := []string{"OVER_UNDER_25", "1X2"}

		st := stream.NewEventMarketStreamer(fixClient, logger, clock, markets)

		ctx := context.Background()

		fixReq := mock.MatchedBy(func(r *statistico.FixtureSearchRequest) bool {
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
				Round:       &statistico.Round{Name: "5"},
			},
		}

		fixClient.On("Search", ctx, fixReq).Return(fixtures, nil)

		event := mock.MatchedBy(func(e *exchange.Event) bool {
			assert.Equal(t, "OVER_UNDER_35", e.Market)
			assert.Equal(t, time.Unix(1547465400, 0), e.Date)
			assert.Equal(t, "West Ham United v Arsenal", e.Name)
			assert.Equal(t, uint64(349811), e.ID)
			return true
		})

		marketOne := exchange.Market{
			ID:       "1.254912",
			EventID:  349811,
			Name:     "OVER_UNDER_35",
			Exchange: "BETFAIR",
			Runners: []*exchange.Runner{
				{
					ID:         "0",
					Name:       "OVER",
					BackPrices: nil,
					LayPrices:  nil,
				},
			},
		}

		factory.On("CreateMarket", ctx, event).Return(&marketOne, nil)

		from := time.Date(2019, 01, 14, 10, 00, 00, 00, time.UTC)
		to := time.Date(2019, 01, 14, 12, 00, 00, 00, time.UTC)

		ch := st.Stream(ctx, from, to, factory, "OVER_UNDER_35")

		emOne := <-ch

		<-ch

		a := assert.New(t)

		a.Equal("1.254912", emOne.ID)
		a.Equal(uint64(349811), emOne.EventID)
		a.Equal(uint64(8), emOne.CompetitionID)
		a.Equal(uint64(17420), emOne.SeasonID)
		a.Equal(uint64(5), emOne.Round)
		a.Equal(int64(1547465400), emOne.EventDate)
		a.Equal("OVER_UNDER_35", emOne.MarketName)
		a.Equal("BETFAIR", emOne.Exchange)
		a.Equal(marketOne.Runners, emOne.Runners)
		a.Equal(int64(1547465100), emOne.Timestamp)

		fixClient.AssertExpectations(t)
		factory.AssertExpectations(t)
	})

	t.Run("factory is not called if event date and current date difference is less than zero", func(t *testing.T) {
		t.Helper()

		fixClient := new(MockFixtureClient)
		factory := new(exchange.MockMarketFactory)
		logger, _ := test.NewNullLogger()
		clock := clockwork.NewFakeClockAt(time.Date(2019, 01, 14, 20, 25, 00, 00, time.UTC))
		markets := []string{"OVER_UNDER_25", "1X2"}

		st := stream.NewEventMarketStreamer(fixClient, logger, clock, markets)

		fixReq := mock.MatchedBy(func(r *statistico.FixtureSearchRequest) bool {
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

		factory.AssertNotCalled(t, "CreateMarket")

		ch := st.Stream(ctx, from, to, factory, "")

		one := <-ch

		assert.Nil(t, one)
	})

	t.Run("returns nil if error fetching fixtures via event client", func(t *testing.T) {
		t.Helper()

		fixClient := new(MockFixtureClient)
		factory := new(exchange.MockMarketFactory)
		logger, hook := test.NewNullLogger()
		clock := clockwork.NewFakeClockAt(time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC))
		markets := []string{"OVER_UNDER_25", "1X2"}

		st := stream.NewEventMarketStreamer(fixClient, logger, clock, markets)

		fixReq := mock.MatchedBy(func(r *statistico.FixtureSearchRequest) bool {
			assert.Equal(t, "2019-01-14T10:00:00Z", r.DateAfter.GetValue())
			assert.Equal(t, "2019-01-14T12:00:00Z", r.DateBefore.GetValue())
			return true
		})

		ctx := context.Background()

		fixClient.On("Search", ctx, fixReq).Return([]*statistico.Fixture{}, errors.New("error fetching fixtures"))

		from := time.Date(2019, 01, 14, 10, 00, 00, 00, time.UTC)
		to := time.Date(2019, 01, 14, 12, 00, 00, 00, time.UTC)

		ch := st.Stream(ctx, from, to, factory, "")

		assert.Nil(t, ch)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "error \"error fetching fixtures\" fetching fixtures in football market requester", hook.LastEntry().Message)

		factory.AssertNotCalled(t, "CreateMarket")
	})

	t.Run("logs info if factory returns an exchange.NoEventMarketError", func(t *testing.T) {
		t.Helper()

		fixClient := new(MockFixtureClient)
		factory := new(exchange.MockMarketFactory)
		logger, hook := test.NewNullLogger()
		clock := clockwork.NewFakeClockAt(time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC))
		markets := []string{"OVER_UNDER_25"}

		st := stream.NewEventMarketStreamer(fixClient, logger, clock, markets)

		ctx := context.Background()

		fixReq := mock.MatchedBy(func(r *statistico.FixtureSearchRequest) bool {
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

		factory.On("CreateMarket", ctx, event).Return(&exchange.Market{}, &exchange.NoEventMarketError{
			Exchange: "BETFAIR",
			Market:   "OVER_UNDER_25",
			EventID:  349811,
		})

		from := time.Date(2019, 01, 14, 10, 00, 00, 00, time.UTC)
		to := time.Date(2019, 01, 14, 12, 00, 00, 00, time.UTC)

		ch := st.Stream(ctx, from, to, factory, "")

		emOne := <-ch

		<-ch

		a := assert.New(t)

		a.Nil(emOne)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
		assert.Equal(t, "No markets returned for event 349811 and market OVER_UNDER_25 and exchange BETFAIR", hook.LastEntry().Message)

		fixClient.AssertExpectations(t)
		factory.AssertExpectations(t)
	})

	t.Run("logs error if factory returns a non exchange.NoEventMarketError error", func(t *testing.T) {
		t.Helper()

		fixClient := new(MockFixtureClient)
		factory := new(exchange.MockMarketFactory)
		logger, hook := test.NewNullLogger()
		clock := clockwork.NewFakeClockAt(time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC))
		markets := []string{"OVER_UNDER_25"}

		st := stream.NewEventMarketStreamer(fixClient, logger, clock, markets)

		ctx := context.Background()

		fixReq := mock.MatchedBy(func(r *statistico.FixtureSearchRequest) bool {
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

		factory.On("CreateMarket", ctx, event).Return(&exchange.Market{}, errors.New("oh no"))

		from := time.Date(2019, 01, 14, 10, 00, 00, 00, time.UTC)
		to := time.Date(2019, 01, 14, 12, 00, 00, 00, time.UTC)

		ch := st.Stream(ctx, from, to, factory, "")

		emOne := <-ch

		<-ch

		a := assert.New(t)

		a.Nil(emOne)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "error when calling factory 'oh no' for event 349811 and market OVER_UNDER_25 and exchange MOCK_FACTORY", hook.LastEntry().Message)

		fixClient.AssertExpectations(t)
		factory.AssertExpectations(t)
	})

	t.Run("process exits if returned market does not contain runners", func(t *testing.T) {
		t.Helper()

		fixClient := new(MockFixtureClient)
		factory := new(exchange.MockMarketFactory)
		logger, _ := test.NewNullLogger()
		clock := clockwork.NewFakeClockAt(time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC))
		markets := []string{"OVER_UNDER_25"}

		st := stream.NewEventMarketStreamer(fixClient, logger, clock, markets)

		ctx := context.Background()

		fixReq := mock.MatchedBy(func(r *statistico.FixtureSearchRequest) bool {
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

		marketOne := exchange.Market{
			ID:       "1.254912",
			EventID:  349811,
			Name:     "OVER_UNDER_25",
			Exchange: "BETFAIR",
			Runners:  []*exchange.Runner{},
		}

		factory.On("CreateMarket", ctx, event).Return(&marketOne, nil)

		from := time.Date(2019, 01, 14, 10, 00, 00, 00, time.UTC)
		to := time.Date(2019, 01, 14, 12, 00, 00, 00, time.UTC)

		ch := st.Stream(ctx, from, to, factory, "")

		emOne := <-ch

		<-ch

		a := assert.New(t)

		a.Nil(emOne)

		fixClient.AssertExpectations(t)
		factory.AssertExpectations(t)
	})
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
