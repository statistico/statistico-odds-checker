package exchange_test

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMarketBuilder_Build(t *testing.T) {
	t.Run("returns a channel of Market struct", func(t *testing.T) {
		t.Helper()

		factoryOne := new(exchange.MockMarketFactory)
		factoryTwo := new(exchange.MockMarketFactory)
		factories := []exchange.MarketFactory{factoryOne, factoryTwo}
		logger, hook := test.NewNullLogger()

		builder := exchange.NewMarketBuilder(factories, logger)

		date, err := time.Parse(time.RFC3339, "2020-03-12T00:00:00+00:00")

		if err != nil {
			t.Fatalf("Error parsing date %s", err.Error())
		}

		event := exchange.Event{
			Date:   date,
			Name:   "West Ham United vs Chelsea",
			ID:     1278121,
			Market: "OVER_UNDER_25",
		}

		ctx := context.Background()

		mkOne := bookmakerMarket("1.5670", "PINNACLE")
		mkTwo := bookmakerMarket("1.2421", "BETFAIR")

		factoryOne.On("CreateMarket", ctx, &event).Once().Return(&mkOne, nil)
		factoryTwo.On("CreateMarket", ctx, &event).Once().Return(&mkTwo, nil)

		markets := builder.Build(ctx, &event)

		one := <-markets
		two := <-markets

		expectedRunners := []*exchange.Runner{
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
		}

		a := assert.New(t)
		a.Equal("1.2421", one.ID)
		a.Equal(uint64(1278121), one.EventID)
		a.Equal("BETFAIR", one.Exchange)
		a.Equal("OVER_UNDER_25", one.Name)
		a.Equal(expectedRunners, one.Runners)
		a.Equal("1.5670", two.ID)
		a.Equal(uint64(1278121), two.EventID)
		a.Equal("PINNACLE", two.Exchange)
		a.Equal("OVER_UNDER_25", one.Name)
		a.Equal(expectedRunners, one.Runners)
		a.Nil(hook.LastEntry())
		factoryOne.AssertExpectations(t)
		factoryTwo.AssertExpectations(t)
	})

	t.Run("logs error if error returned when creating market via market factory implementation", func(t *testing.T) {
		t.Helper()

		factoryOne := new(exchange.MockMarketFactory)
		factoryTwo := new(exchange.MockMarketFactory)
		factories := []exchange.MarketFactory{factoryOne, factoryTwo}
		logger, hook := test.NewNullLogger()

		builder := exchange.NewMarketBuilder(factories, logger)

		date, err := time.Parse(time.RFC3339, "2020-03-12T00:00:00+00:00")

		if err != nil {
			t.Fatalf("Error parsing date %s", err.Error())
		}

		event := exchange.Event{
			Date:   date,
			Name:   "West Ham United vs Chelsea",
			ID:     1278121,
			Market: "OVER_UNDER_25",
		}

		ctx := context.Background()

		mkOne := bookmakerMarket("1.5670", "PINNACLE")

		factoryOne.On("CreateMarket", ctx, &event).Once().Return(&mkOne, nil)
		factoryTwo.On("CreateMarket", ctx, &event).Once().Return(&exchange.Market{}, errors.New("error occurred"))

		markets := builder.Build(ctx, &event)

		<-markets

		assert.Equal(t, 0, len(markets))
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
		assert.Equal(t, "Error when calling client 'error occurred' for event 1278121 and market OVER_UNDER_25", hook.LastEntry().Message)
	})
}

func bookmakerMarket(marketId, ex string) exchange.Market {
	return exchange.Market{
		ID:       marketId,
		Name:     "OVER_UNDER_25",
		EventID:  uint64(1278121),
		Exchange: ex,
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
