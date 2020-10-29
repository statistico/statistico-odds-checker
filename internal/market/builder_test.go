package market_test

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-odds-checker/internal/exchange"
	"github.com/statistico/statistico-odds-checker/internal/exchange/betfair"
	"github.com/statistico/statistico-odds-checker/internal/grpc"
	"github.com/statistico/statistico-odds-checker/internal/grpc/proto"
	"github.com/statistico/statistico-odds-checker/internal/market"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestBuilder_Build(t *testing.T) {
	t.Run("returns a channel of Market struct", func(t *testing.T) {
		t.Helper()

		mr := new(betfair.MockMarketRequester)
		oc := new(grpc.MockOddsCompilerGrpcClient)
		logger, hook := test.NewNullLogger()

		builder := market.NewBuilder(oc, mr, logger)

		date, err := time.Parse(time.RFC3339, "2020-03-12T00:00:00+00:00")

		if err != nil {
			t.Fatalf("Error parsing date %s", err.Error())
		}

		query := market.BuilderQuery{
			Date:    date,
			Event:   "West Ham United vs Chelsea",
			EventID: 1278121,
			Sport:   "football",
			Markets: []string{"OVER_UNDER_25"},
		}

		ctx := context.Background()

		odds := protoOdds(1.25, 11)

		oc.On("GetEventMarket", ctx, uint64(1278121), "OVER_UNDER_25").Return(odds, nil)

		exQuery := mock.MatchedBy(func(r *exchange.Query) bool {
			assert.Equal(t, "West Ham United vs Chelsea", r.Event)
			assert.Equal(t, "football", r.Sport)
			assert.Equal(t, date, r.Date)
			assert.Equal(t, "OVER_UNDER_25", r.Market)
			return true
		})

		mk := bookmakerMarket("1.2421")

		mr.On("Fetch", ctx, exQuery).Once().Return(mk, nil)

		markets := builder.Build(ctx, &query)

		one := <-markets

		a := assert.New(t)
		a.Equal(uint64(1278121), one.EventID)
		a.Equal(odds, one.ImpliedOdds)
		a.Equal(mk, &one.ExchangeMarket)
		assert.Nil(t, hook.LastEntry())
		mr.AssertExpectations(t)
		oc.AssertExpectations(t)
	})

	t.Run("logs error if error returned when fetching odds compiler event market", func(t *testing.T) {
		t.Helper()

		mr := new(betfair.MockMarketRequester)
		oc := new(grpc.MockOddsCompilerGrpcClient)
		logger, hook := test.NewNullLogger()

		builder := market.NewBuilder(oc, mr, logger)

		date, err := time.Parse(time.RFC3339, "2020-03-12T00:00:00+00:00")

		if err != nil {
			t.Fatalf("Error parsing date %s", err.Error())
		}

		query := market.BuilderQuery{
			Date:    date,
			Event:   "West Ham United vs Chelsea",
			EventID: 1278121,
			Sport:   "football",
			Markets: []string{"OVER_UNDER_25"},
		}

		ctx := context.Background()

		oc.On("GetEventMarket", ctx, uint64(1278121), "OVER_UNDER_25").Return([]*proto.Odds{}, errors.New("error occurred"))

		mr.AssertNotCalled(t, "Fetch")

		markets := builder.Build(ctx, &query)

		<-markets

		assert.Equal(t, 0, len(markets))
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Error when calling client 'error occurred' for event 1278121 and market OVER_UNDER_25", hook.LastEntry().Message)
	})

	t.Run("logs error if error returned when fetching market via market requester", func(t *testing.T) {
		t.Helper()

		mr := new(betfair.MockMarketRequester)
		oc := new(grpc.MockOddsCompilerGrpcClient)
		logger, hook := test.NewNullLogger()

		builder := market.NewBuilder(oc, mr, logger)

		date, err := time.Parse(time.RFC3339, "2020-03-12T00:00:00+00:00")

		if err != nil {
			t.Fatalf("Error parsing date %s", err.Error())
		}

		query := market.BuilderQuery{
			Date:    date,
			Event:   "West Ham United vs Chelsea",
			EventID: 1278121,
			Sport:   "football",
			Markets: []string{"OVER_UNDER_25"},
		}

		ctx := context.Background()

		odds := protoOdds(1.25, 11)

		oc.On("GetEventMarket", ctx, uint64(1278121), "OVER_UNDER_25").Return(odds, nil)

		exQuery := mock.MatchedBy(func(r *exchange.Query) bool {
			assert.Equal(t, "West Ham United vs Chelsea", r.Event)
			assert.Equal(t, "football", r.Sport)
			assert.Equal(t, date, r.Date)
			assert.Equal(t, "OVER_UNDER_25", r.Market)
			return true
		})

		mr.On("Fetch", ctx, exQuery).Once().Return(&exchange.Market{}, errors.New("error occurred"))

		markets := builder.Build(ctx, &query)

		<-markets

		assert.Equal(t, 0, len(markets))
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Error when calling client 'error occurred' for event 1278121 and market OVER_UNDER_25", hook.LastEntry().Message)
	})
}

func protoOdds(over, under float32) []*proto.Odds {
	return []*proto.Odds{
		{
			Selection: "over",
			Price:     over,
		},
		{
			Selection: "over",
			Price:     under,
		},
	}
}

func bookmakerMarket(marketId string) *exchange.Market {
	return &exchange.Market{ID: marketId}
}
