package market_test

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange/betfair"
	"github.com/statistico/statistico-odds-checker/internal/app/market"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestBuilder_Build(t *testing.T) {
	t.Run("returns a channel of Market struct", func(t *testing.T) {
		t.Helper()

		mr := new(betfair.MockMarketFactory)
		logger, hook := test.NewNullLogger()

		builder := market.NewBuilder(mr, logger)

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

		exQuery := mock.MatchedBy(func(r *exchange.Event) bool {
			assert.Equal(t, "West Ham United vs Chelsea", r.Name)
			assert.Equal(t, date, r.Date)
			assert.Equal(t, "OVER_UNDER_25", r.Market)
			return true
		})

		mk := bookmakerMarket("1.2421")

		mr.On("CreateMarket", ctx, exQuery).Once().Return(mk, nil)

		markets := builder.Build(ctx, &query)

		one := <-markets

		a := assert.New(t)
		a.Equal("1.2421", one.ID)
		a.Equal(uint64(1278121), one.EventID)
		a.Equal("betfair", one.Exchange)
		a.Equal("OVER_UNDER_25", one.Name)
		a.Equal(mk.Runners, one.ExchangeRunners)
		assert.Nil(t, hook.LastEntry())
		mr.AssertExpectations(t)
	})

	t.Run("logs error if error returned when fetching market via market requester", func(t *testing.T) {
		t.Helper()

		mr := new(betfair.MockMarketFactory)
		logger, hook := test.NewNullLogger()

		builder := market.NewBuilder(mr, logger)

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

		exQuery := mock.MatchedBy(func(r *exchange.Event) bool {
			assert.Equal(t, "West Ham United vs Chelsea", r.Name)
			assert.Equal(t, date, r.Date)
			assert.Equal(t, "OVER_UNDER_25", r.Market)
			return true
		})

		mr.On("CreateMarket", ctx, exQuery).Once().Return(&exchange.Market{}, errors.New("error occurred"))

		markets := builder.Build(ctx, &query)

		<-markets

		assert.Equal(t, 0, len(markets))
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Error when calling client 'error occurred' for event 1278121 and market OVER_UNDER_25", hook.LastEntry().Message)
	})
}

func bookmakerMarket(marketId string) *exchange.Market {
	return &exchange.Market{
		ID:           marketId,
		Name:         "OVER_UNDER_25",
		ExchangeName: "betfair",
		Runners: []*exchange.Runner{
			{
				ID:   49792,
				Name: "Over 2.5 Goals",
				BackPrices: []exchange.PriceSize{
					{
						Price: 1.54,
						Size:  1301.00,
					},
				},
			},
		},
	}
}
