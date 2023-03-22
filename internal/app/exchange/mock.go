package exchange

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MockMarketBuilder struct {
	mock.Mock
}

func (m *MockMarketBuilder) Build(ctx context.Context, e *Event) <-chan *Market {
	args := m.Called(ctx, e)
	return args.Get(0).(<-chan *Market)
}

type MockMarketFactory struct {
	mock.Mock
}

func (m *MockMarketFactory) CreateMarket(ctx context.Context, e *Event) (*Market, error) {
	args := m.Called(ctx, e)
	return args.Get(0).(*Market), args.Error(1)
}

func (m *MockMarketFactory) Exchange() string {
	return "MOCK_FACTORY"
}
