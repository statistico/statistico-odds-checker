package sportmonks

import (
	"context"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/stretchr/testify/mock"
)

type MockOddsParser struct {
	mock.Mock
}

func (m *MockOddsParser) ParseMarketRunners(ctx context.Context, fixtureID, exchangeID int, market string) ([]*exchange.Runner, error) {
	args := m.Called(ctx, fixtureID, exchangeID, market)
	return args.Get(0).([]*exchange.Runner), args.Error(1)
}
