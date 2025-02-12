package statistico_test

import (
	"context"
	"errors"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange/statistico"
	sp "github.com/statistico/statistico-proto/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestMarketFactory_CreateMarket(t *testing.T) {
	t.Run("returns a built market for an event", func(t *testing.T) {
		t.Helper()

		compiler := new(MockOddsClient)
		factory := statistico.NewMarketFactory(compiler)

		ctx := context.Background()

		event := exchange.Event{
			ID:     55,
			Name:   "West Ham vs Chelsea",
			Date:   time.Time{},
			Market: "OVER_UNDER_25",
		}

		em := sp.EventMarket{
			EventId: 55,
			Market:  "OVER_UNDER_25",
			Odds: []*sp.Odds{
				{
					Price:     1.54,
					Selection: "OVER",
				},
				{
					Price:     2.34,
					Selection: "UNDER",
				},
			},
		}

		compiler.On("GetEventMarket", ctx, event.ID, event.Market).Return(&em, nil)

		expectedMarket := &exchange.Market{
			ID:       "STA-55-OVER_UNDER_25",
			Exchange: "STATISTICO",
			Name:     "OVER_UNDER_25",
			EventID:  55,
			Runners: []*exchange.Runner{
				{
					ID:    0,
					Label: "OVER",
					BackPrices: []exchange.PriceSize{
						{
							Price: 1.54,
							Size:  0,
						},
					},
				},
				{
					ID:    0,
					Label: "UNDER",
					BackPrices: []exchange.PriceSize{
						{
							Price: 2.34,
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
		compiler.AssertExpectations(t)
	})

	t.Run("returns an error if error is returned by compiler client", func(t *testing.T) {
		t.Helper()

		compiler := new(MockOddsClient)
		factory := statistico.NewMarketFactory(compiler)

		ctx := context.Background()

		event := exchange.Event{
			ID:     55,
			Name:   "West Ham vs Chelsea",
			Date:   time.Time{},
			Market: "OVER_UNDER_25",
		}

		compiler.On("GetEventMarket", ctx, event.ID, event.Market).Return(&sp.EventMarket{}, errors.New("oh no"))

		_, err := factory.CreateMarket(ctx, &event)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})

	t.Run("return a nil market if market is not over under 2.5 goals", func(t *testing.T) {
		t.Helper()

		compiler := new(MockOddsClient)
		factory := statistico.NewMarketFactory(compiler)

		ctx := context.Background()

		event := exchange.Event{
			ID:     55,
			Name:   "West Ham vs Chelsea",
			Date:   time.Time{},
			Market: "OVER_UNDER_45",
		}

		compiler.AssertNotCalled(t, "GetEventMarket")

		em, err := factory.CreateMarket(ctx, &event)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Nil(t, em)
	})
}

type MockOddsClient struct {
	mock.Mock
}

func (m *MockOddsClient) GetEventMarket(ctx context.Context, eventID uint64, market string) (*sp.EventMarket, error) {
	args := m.Called(ctx, eventID, market)
	return args.Get(0).(*sp.EventMarket), args.Error(1)
}
