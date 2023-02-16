package betcris_test

import (
	"context"
	"errors"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange/betcris"
	"github.com/statistico/statistico-odds-checker/internal/app/sportmonks"
	sp "github.com/statistico/statistico-sportmonks-go-client"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMarketFactory_CreateMarket(t *testing.T) {
	t.Run("returns a built market for an event", func(t *testing.T) {
		t.Helper()

		parser := new(sportmonks.MockOddsParser)

		factory := betcris.NewMarketFactory(parser)

		ctx := context.Background()

		event := exchange.Event{
			ID:     55,
			Name:   "BETCRIS",
			Date:   time.Time{},
			Market: "OVER_UNDER_25",
		}

		odds := []sp.Odds{
			{
				Label: "Over",
				Dp3:   "1.95",
			},
			{
				Label: "Under",
				Dp3:   "2.08",
			},
		}

		parser.On("ParseMarketOdds", ctx, 55, 13, "OVER_UNDER_25").Return(odds, nil)

		expectedMarket := &exchange.Market{
			ID:       "BETCRIS-55-OVER_UNDER_25",
			Exchange: "BETCRIS",
			Name:     "OVER_UNDER_25",
			EventID:  55,
			Runners: []*exchange.Runner{
				{
					ID:   0,
					Name: "OVER",
					BackPrices: []exchange.PriceSize{
						{
							Price: 1.95,
							Size:  0,
						},
					},
				},
				{
					ID:   0,
					Name: "UNDER",
					BackPrices: []exchange.PriceSize{
						{
							Price: 2.08,
							Size:  0,
						},
					},
				},
			},
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

		factory := betcris.NewMarketFactory(parser)

		ctx := context.Background()

		event := exchange.Event{
			ID:     55,
			Name:   "BETCRIS",
			Date:   time.Time{},
			Market: "OVER_UNDER_25",
		}

		parser.On("ParseMarketOdds", ctx, 55, 13, "OVER_UNDER_25").Return([]sp.Odds{}, errors.New("error from sportmonks"))

		_, err := factory.CreateMarket(ctx, &event)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})

	t.Run("returns an error if odds slice returned by odds parser is empty", func(t *testing.T) {
		t.Helper()

		parser := new(sportmonks.MockOddsParser)

		factory := betcris.NewMarketFactory(parser)

		ctx := context.Background()

		event := exchange.Event{
			ID:     55,
			Name:   "BETCRIS",
			Date:   time.Time{},
			Market: "OVER_UNDER_25",
		}

		parser.On("ParseMarketOdds", ctx, 55, 13, "OVER_UNDER_25").Return([]sp.Odds{}, nil)

		_, err := factory.CreateMarket(ctx, &event)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})

	t.Run("returns an error if unable to parse odds for a runner", func(t *testing.T) {
		t.Helper()

		parser := new(sportmonks.MockOddsParser)

		factory := betcris.NewMarketFactory(parser)

		ctx := context.Background()

		event := exchange.Event{
			ID:     55,
			Name:   "BETCRIS",
			Date:   time.Time{},
			Market: "OVER_UNDER_25",
		}

		odds := []sp.Odds{
			{
				Label: "Over",
				Dp3:   "Hello",
			},
			{
				Label: "Under",
				Dp3:   "2.08",
			},
		}

		parser.On("ParseMarketOdds", ctx, 55, 13, "OVER_UNDER_25").Return(odds, nil)

		_, err := factory.CreateMarket(ctx, &event)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})
}
