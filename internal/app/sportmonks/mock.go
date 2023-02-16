package sportmonks

import (
	"context"
	sportmonks "github.com/statistico/statistico-sportmonks-go-client"
	"github.com/stretchr/testify/mock"
)

type MockOddsParser struct {
	mock.Mock
}

func (m *MockOddsParser) ParseMarketOdds(ctx context.Context, fixtureID, exchangeID int, market string) ([]sportmonks.Odds, error) {
	args := m.Called(ctx, fixtureID, exchangeID, market)
	return args.Get(0).([]sportmonks.Odds), args.Error(1)
}
