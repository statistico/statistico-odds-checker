package betfair

import (
	"context"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/stretchr/testify/mock"
)

type MockMarketFactory struct {
	mock.Mock
}

func (m *MockMarketFactory) CreateMarket(ctx context.Context, e *exchange.Event) (*exchange.Market, error) {
	args := m.Called(ctx, e)
	return args.Get(0).(*exchange.Market), args.Error(1)
}
