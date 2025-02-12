package bet365_test

import (
	"context"
	"errors"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange/bet365"
	"github.com/statistico/statistico-odds-checker/internal/app/sportmonks"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMarketFactory_CreateMarket(t *testing.T) {
	t.Run("returns a built market for an event", func(t *testing.T) {
		t.Helper()

		parser := new(sportmonks.MockOddsParser)

		factory := bet365.NewMarketFactory(parser)

		ctx := context.Background()

		event := exchange.Event{
			ID:     55,
			Name:   "BET365",
			Date:   time.Time{},
			Market: "MATCH_ODDS",
		}

		runners := []*exchange.Runner{
			{
				ID:   0,
				Name: "HOME",
				BackPrices: []exchange.PriceSize{
					{
						Price: 1.95,
						Size:  0,
					},
				},
			},
			{
				ID:   0,
				Name: "AWAY",
				BackPrices: []exchange.PriceSize{
					{
						Price: 2.08,
						Size:  0,
					},
				},
			},
			{
				ID:   0,
				Name: "DRAW",
				BackPrices: []exchange.PriceSize{
					{
						Price: 2.75,
						Size:  0,
					},
				},
			},
		}

		parser.On("ParseMarketRunners", ctx, 55, 2, "MATCH_ODDS").Return(runners, nil)

		expectedMarket := &exchange.Market{
			ID:       "BET365-55-MATCH_ODDS",
			Exchange: "BET365",
			Name:     "MATCH_ODDS",
			EventID:  55,
			Runners:  runners,
		}

		market, err := factory.CreateMarket(ctx, &event)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, expectedMarket, market)
		parser.AssertExpectations(t)
	})

	t.Run("returns an error if error returned by odds parser", func(t *testing.T) {
		t.Helper()

		parser := new(sportmonks.MockOddsParser)

		factory := bet365.NewMarketFactory(parser)

		ctx := context.Background()

		event := exchange.Event{
			ID:     55,
			Name:   "BET365",
			Date:   time.Time{},
			Market: "OVER_UNDER_25",
		}

		parser.On("ParseMarketRunners", ctx, 55, 2, "OVER_UNDER_25").Return([]*exchange.Runner{}, errors.New("error from sportmonks"))

		_, err := factory.CreateMarket(ctx, &event)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})

	t.Run("returns an error if odds slice returned by odds parser is empty", func(t *testing.T) {
		t.Helper()

		parser := new(sportmonks.MockOddsParser)

		factory := bet365.NewMarketFactory(parser)

		ctx := context.Background()

		event := exchange.Event{
			ID:     55,
			Name:   "BET365",
			Date:   time.Time{},
			Market: "OVER_UNDER_25",
		}

		parser.On("ParseMarketRunners", ctx, 55, 2, "OVER_UNDER_25").Return([]*exchange.Runner{}, nil)

		_, err := factory.CreateMarket(ctx, &event)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})
}
