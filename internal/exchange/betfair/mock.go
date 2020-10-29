package betfair

import (
	"context"
	"github.com/statistico/statistico-odds-checker/internal/exchange"
	"github.com/stretchr/testify/mock"
)

type MockMarketRequester struct {
	mock.Mock
}

func (m *MockMarketRequester) Fetch(ctx context.Context, q *exchange.Query) (*exchange.Market, error) {
	args := m.Called(ctx, q)
	return args.Get(0).(*exchange.Market), args.Error(1)
}
